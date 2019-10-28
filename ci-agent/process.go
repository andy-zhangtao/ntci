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
		case s := <-status:
			switch s.Stauts {
			case store.BuildSuccess:
				d, err := bus.Pb.GetBuildByID(s.User, s.Name, s.Id)
				if err != nil {
					logrus.Errorf("Query Build Error: %s. Filter: user:%s. name: %s. id: %d", err, s.User, s.Name, s.Id)
					bus.Pb.UpdataBuildStatus(store.ProcessFailed, s.Id, s.Name, s.User)
					break
				}

				nts, err := bus.Pb.GetNtci(s.User, s.Name, d.Branch)
				if err != nil {
					logrus.Errorf("Get Ntci Error: %s. Filter: user:%s. name: %s. id: %d", err, s.User, s.Name, s.Id)
					bus.Pb.UpdataBuildStatus(store.ProcessFailed, s.Id, s.Name, s.User)
					break
				}

				var nt git.Ntci

				err = yaml.Unmarshal([]byte(nts), &nt)
				if err != nil {
					logrus.Errorf("Unmarshal Ntci Error: %s. Content: %s", err, nts)
					bus.Pb.UpdataBuildStatus(store.ProcessFailed, s.Id, s.Name, s.User)
					break
				}

				if len(nt.Deployer) > 0 {
					for filter, value := range nt.Deployer {
						if addr, ok := bus.Deployer[filter]; ok {
							params, err := yaml.Marshal(value)
							if err != nil {
								logrus.Errorf("Marshal Ntci Error: %s. Content: %s", err, nt.Deployer)
								bus.Pb.UpdataBuildStatus(store.ProcessFailed, s.Id, s.Name, s.User)
								break
							}
							err = invokeDeployer(addr, string(params))
							if err != nil {
								logrus.Errorf("Invoke Deployer Error: %s. ", err)
								bus.Pb.UpdataBuildStatus(store.ProcessFailed, s.Id, s.Name, s.User)
								break
							}
						}
					}
				}
			}
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
