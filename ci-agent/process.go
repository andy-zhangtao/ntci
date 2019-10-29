package main

import (
	"context"
	"errors"
	"fmt"

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
			logrus.Infof("Receive status update msg")
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
		bus.Pb.UpdataBuildStatus(store.ProcessFailed, id, user, name)
		return
	}

	nts, err := bus.Pb.GetNtci(user, name, d.Branch)
	if err != nil {
		logrus.Errorf("Get Ntci Error: %s. Filter: user:%s. name: %s. id: %d", err, user, name, id)
		bus.Pb.UpdataBuildStatus(store.ProcessFailed, id, user, name)
		return
	}

	var nt git.Ntci

	err = yaml.Unmarshal([]byte(nts), &nt)
	if err != nil {
		logrus.Errorf("Unmarshal Ntci Error: %s. Content: %s", err, nts)
		bus.Pb.UpdataBuildStatus(store.ProcessFailed, id, user, name)
		return
	}

	logrus.Infof("Deploy: [%v]", nt.Deployer)
	if len(nt.Deployer) > 0 {
		for filter, value := range nt.Deployer {
			if addr, ok := bus.Deployer[filter]; ok {
				params, err := yaml.Marshal(value)
				if err != nil {
					logrus.Errorf("Marshal Ntci Error: %s. Content: %s", err, nt.Deployer)
					bus.Pb.UpdataBuildStatus(store.ProcessFailed, id, user, name)
					return
				}

				logrus.Infof("k8s name: %s addr: %s params: %s", filter, addr, string(params))
				err = invokeDeployer(addr, string(params))
				if err != nil {
					logrus.Errorf("Invoke Deployer Error: %s. ", err)
					bus.Pb.UpdataBuildStatus(store.ProcessFailed, id, user, name)
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
