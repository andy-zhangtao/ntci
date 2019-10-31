package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"ntci/ci-agent/dataBus"
	"ntci/ci-agent/git"
	"ntci/ci-agent/store"
	deploy_rpc_v1 "ntci/ci-grpc/deploy"
)

func control() {
	logrus.Infof("Process Control Start")
	bus := dataBus.GetBus()

	for {
		select {
		case s := <-bus.JobStatus:
			switch s.Stauts {
			case store.BuildSuccess:
				go deploy(s.User, s.Name, s.Id)
			}
		}
	}
}

func deploy(user, name string, id int) {
	bus := dataBus.GetBus()
	d, err := bus.Pb.GetBuildByID(user, name, id)
	if err != nil {
		logrus.Errorf("Query Build Error: %s. Filter: user:%s. name: %s. id: %d", err, user, name, id)
		bus.Pb.UpdataBuildStatus(store.DeployFailed, id, user, name)
		return
	}

	nts, err := bus.Pb.GetNtci(user, name, d.Branch)
	if err != nil {
		logrus.Errorf("Get Ntci Error: %s. Filter: user:%s. name: %s. id: %d", err, user, name, id)
		bus.Pb.UpdataBuildStatus(store.DeployFailed, id, user, name)
		return
	}

	var nt git.Ntci

	err = yaml.Unmarshal([]byte(nts), &nt)
	if err != nil {
		logrus.Errorf("Unmarshal Ntci Error: %s. Content: %s", err, nts)
		bus.Pb.UpdataBuildStatus(store.DeployFailed, id, user, name)
		return
	}

	logrus.Debugf("Deploy: [%v]", nt.Deployer)
	if len(nt.Deployer) > 0 {
		if ploy, ok := nt.Deployer[d.Branch]; ok {
			for filter, value := range ploy {
				if addr, ok := bus.Deployer[filter]; ok {
					params, err := yaml.Marshal(value)
					if err != nil {
						logrus.Errorf("Marshal Ntci Error: %s. Content: %s", err, nt.Deployer)
						err = bus.Pb.UpdataBuildStatus(store.DeployFailed, id, name, user)
						if err != nil {
							logrus.Errorf("Update Deployer Error: %s. ", err)
						}
						return
					}

					p := string(params)
					cp := environmentConver(p, nt, d)
					logrus.Infof("k8s name: %s addr: %s params: %s env conver: %s", filter, addr, p, cp)
					err = invokeDeployer(addr, cp)
					if err != nil {
						logrus.Errorf("Invoke Deployer Error: %s. ", err)
						err = bus.Pb.UpdataBuildStatus(store.DeployFailed, id, name, user)
						if err != nil {
							logrus.Errorf("Update Deployer Error: %s. ", err)
						}
						return
					}
				}
			}

			err = bus.Pb.UpdataBuildStatus(store.DeploySuccess, id, name, user)
			if err != nil {
				logrus.Errorf("Update Deployer Error: %s. ", err)
			}
		}

	}
}

func environmentConver(params string, nt git.Ntci, d store.Build) string {
	if strings.Contains(params, "$") {

		subStr := strings.Split(params, "$")
		result := subStr[0]

		_subStr := subStr[1:]

		env := fillNTCIEnvironment(nt, d)
		for _, s := range _subStr {
			result += converEnv(s, env)
		}
		return result
	}

	return params
}

func converEnv(s string, env map[string]string) string {
	for i := 0; i < len(s)-1; i++ {
		if env[s[0:i+1]] != "" {
			return env[s[0:i+1]] + s[i+1:]
		}
		c := s[i+1]
		if (c < 48 || (c >= 58 && c <= 64) || (c >= 91 && c <= 96) || c >= 123) && (c != 95 && c != 45) {
			return s[i+1:]
		}
	}
	return ""
}

func invokeDeployer(addr, params string) (err error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %v", err))
	}

	defer conn.Close()

	c := deploy_rpc_v1.NewDeployServiceClient(conn)

	reply, err := c.RestartJob(context.Background(), &deploy_rpc_v1.DeployRequest{
		Param: params,
	})

	if err != nil {
		return err
	}

	if reply.Code != 0 {
		return errors.New(reply.Msg)
	}

	return nil
}

func fillNTCIEnvironment(nt git.Ntci, b store.Build) map[string]string {
	var env = map[string]string{
		"NTCI_BUILDER_SHA":    b.Sha,
		"NTCI_BUILDER_USER":   b.User,
		"NTCI_BUILDER_JID":    strconv.Itoa(b.Id),
		"NTCI_BUILDER_GIT":    b.Git,
		"NTCI_BUILDER_BRANCH": b.Branch,
	}

	for _, e := range nt.Env {
		_e := strings.Split(e, "=")
		if len(_e) == 1 {
			env[_e[0]] = ""
		} else if len(_e) == 2 {
			env[_e[0]] = _e[1]
		}
	}

	return env
}
