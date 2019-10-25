package rpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"ntci/ci-agent/dataBus"
	"ntci/ci-agent/rpc/builder"
	"ntci/ci-agent/store"
	build_rpc_v1 "ntci/ci-grpc/build"
	gateway_rpc_v1 "ntci/ci-grpc/gateway"
)

type gateway struct {
	buildAddr string
}

func (g *gateway) RestartJob(ctx context.Context, in *gateway_rpc_v1.Builder) (*gateway_rpc_v1.Reply, error) {
	bus := dataBus.GetBus()

	id, _ := strconv.Atoi(in.Jid)

	build, err := bus.Pb.GetBuildByID(in.User, in.Jname, id)
	if err != nil {
		return &gateway_rpc_v1.Reply{
			Code:    -1,
			Message: err.Error(),
		}, err
	}

	env, err := bus.Pb.GetCommonEnv()
	if err != nil {
		logrus.Error(err)
	}

	r, err := builder.InvokeBuilderService(&build_rpc_v1.Request{
		Name:       build.Name,
		Branch:     build.Branch,
		Url:        build.Git,
		Id:         int32(build.Id),
		Language:   build.Language,
		Lanversion: build.Lanversion,
		User:       build.User,
		Sha:        build.Sha,
		Message:    build.Message,
		Env:        env,
	})

	if err != nil {
		bus.Pb.UpdataBuildStatus(int32(store.BuildFailed), build.Id, build.Name, build.User)
		logrus.Errorf("Invoke Build Service Error.  %v", err)
		return &gateway_rpc_v1.Reply{
			Code:    -1,
			Message: err.Error(),
		}, err
	}

	if r.Code != 0 {
		bus.Pb.UpdataBuildStatus(int32(store.BuildFailed), build.Id, build.Name, build.User)
		logrus.Errorf("Invoke Build Service Failed.  %d, %s", r.Code, r.Message)
		return &gateway_rpc_v1.Reply{
			Code:    -1,
			Message: fmt.Sprintf("Invoke Build Service Failed.  %d, %s", r.Code, r.Message),
		}, errors.New(fmt.Sprintf("Invoke Build Service Failed.  %d, %s", r.Code, r.Message))
	}

	bus.Pb.UpdataBuildStatus(int32(store.BuildEnv), build.Id, build.Name, build.User)

	return &gateway_rpc_v1.Reply{
		Code:    0,
		Message: "OK",
	}, nil
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
