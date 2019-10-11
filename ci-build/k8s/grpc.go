package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	build_rpc_v1 "ntci/ci-grpc/build"
)

type server struct{}

/*
Ping

Health Check
*/
func (s *server) Ping(ctx context.Context, in *build_rpc_v1.Request) (*build_rpc_v1.Reply, error) {

	logrus.Debug("Receive Ping Request")
	return &build_rpc_v1.Reply{
		Code:    0,
		Message: "OK",
	}, nil
}

func (s *server) Run(ctx context.Context, in *build_rpc_v1.Request) (*build_rpc_v1.Reply, error) {

	logrus.Debugf("Receive Build Request. Name: %s Branch: %s Git: %s ID: %s ", in.Name, in.Branch, in.Url, in.Id)

	return &build_rpc_v1.Reply{
		Code:    0,
		Message: "OK",
	}, nil
}

func start(port int) {

	p := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", p)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	build_rpc_v1.RegisterBuildServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
