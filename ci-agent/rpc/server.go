package rpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"ntci/ci-agent/dataBus"
	build_rpc_v1 "ntci/ci-grpc/build"
	gateway_rpc_v1 "ntci/ci-grpc/gateway"
)

type gateway struct {
	buildAddr string
}

func (g *gateway) GetBuild(ctx context.Context, in *gateway_rpc_v1.BuildRequest) (*gateway_rpc_v1.JobInfo, error) {
	conn, err := grpc.Dial(g.buildAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := build_rpc_v1.NewBuildServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	j, err := c.GetJob(ctx, &build_rpc_v1.JobRequest{
		Owner: in.User,
		Name:  in.Name,
	})

	if err != nil {
		logrus.Errorf("Fetch Build Error: %v", err)
		return nil, errors.New(fmt.Sprintf("Fetch Build Error: %v", err))
	}

	result := new(gateway_rpc_v1.JobInfo)
	result.Count = j.Count

	var js []*gateway_rpc_v1.JobDetail
	for _, ji := range j.Jd {
		js = append(js, &gateway_rpc_v1.JobDetail{
			Name:      ji.Name,
			Status:    ji.Status,
			Timestamp: ji.Timestamp,
			Branch:    ji.Branch,
			Url:       ji.Url,
		})
	}

	result.Jd = js
	return result, nil
}

func Run(port int) {
	bus := dataBus.GetBus()

	p := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", p)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	gateway_rpc_v1.RegisterGateWayRpcServer(s, &gateway{
		buildAddr: bus.Build[bus.BuildMode].Addr,
	})

	reflection.Register(s)

	logrus.Infof("GateWay Listen on: %d", port)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
