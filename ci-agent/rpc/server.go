package rpc

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"ntci/ci-agent/dataBus"
	gateway_rpc_v1 "ntci/ci-grpc/gateway"
)

type gateway struct {
	buildAddr string
}

func (g *gateway) JobStatus(ctx context.Context, in *gateway_rpc_v1.Builder) (*gateway_rpc_v1.Reply, error) {
	bus := dataBus.GetBus()
	id, _ := strconv.Atoi(in.Jid)

	err := bus.Pb.UpdataBuildStatus(in.Status, id, in.Jname, in.User)

	if err != nil {
		return &gateway_rpc_v1.Reply{
			Code:    -1,
			Message: err.Error(),
		}, nil
	}

	return &gateway_rpc_v1.Reply{
		Code:    0,
		Message: "OK",
	}, nil
}

func (g *gateway) GetBuild(ctx context.Context, in *gateway_rpc_v1.BuildRequest) (*gateway_rpc_v1.JobInfo, error) {

	result := new(gateway_rpc_v1.JobInfo)

	bus := dataBus.GetBus()
	bs, err := bus.Pb.GetBuild(in.User, in.Name)
	if err != nil {
		return nil, err
	}

	var js []*gateway_rpc_v1.JobDetail
	for _, ji := range bs {
		js = append(js, &gateway_rpc_v1.JobDetail{
			Name:      ji.Name,
			Status:    int32(ji.Status),
			Timestamp: ji.Timestamp.Format("2006-01-02 15:04:05"),
			Branch:    ji.Branch,
			Url:       ji.Git,
			Id:        int32(ji.Id),
			Sha:       ji.Sha,
			Message:   ji.Message,
		})
	}

	result.Jd = js
	result.Count = int32(len(bs))

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
