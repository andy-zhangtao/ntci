package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"ntci/ci-build/k8s/deploy"
	"ntci/ci-build/k8s/store"
	build_rpc_v1 "ntci/ci-grpc/build"
)

type server struct {
	pg *store.PGBus
}

/*
fetch Image.

If this build server don't contains specify language and version, then return false.

Otherwise return true and fully image name.
*/
func fetchImage(language, version string) (bool, string) {

	image := ""
	if _, ok := bus.LanguageRuntime[language]; !ok {
		return false, image
	}

	l := bus.LanguageRuntime[language]

	if _, ok := l[version]; !ok {
		return false, image
	}

	return true, fmt.Sprintf("%s:%s", l[version], version)
}

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

/*
Run

Create Build Environment In K8s Cluster.

Run() will store build info into db.
*/
func (s *server) Run(ctx context.Context, in *build_rpc_v1.Request) (*build_rpc_v1.Reply, error) {

	logrus.Debugf("Receive Build Request. Name: %s Branch: %s Git: %s ID: %s ", in.Name, in.Branch, in.Url, in.Id)

	b := store.Build{
		Name:      in.Name,
		Branch:    in.Branch,
		Git:       in.Url,
		Timestamp: time.Now(),
	}

	isExist, image := fetchImage(in.Language, in.Lanversion)
	if !isExist {
		logrus.Errorf("Can not support this language: %s %s", in.Language, in.Lanversion)
		return &build_rpc_v1.Reply{
			Code:    -1,
			Message: fmt.Sprintf("Can not support this language: %s %s", in.Language, in.Lanversion),
		}, nil
	}

	b.Image = image

	id, err := s.pg.AddNewBuild(b)
	if err != nil {
		logrus.Errorf("Add Build Record Error: %s", err.Error())
		return &build_rpc_v1.Reply{
			Code:    -1,
			Message: err.Error(),
		}, nil
	}

	b.Id = id

	err = deploy.NewJob(b)
	if err != nil {
		logrus.Errorf("Create Build Job Error: %s", err.Error())
		return &build_rpc_v1.Reply{
			Code:    -1,
			Message: err.Error(),
		}, nil
	}

	return &build_rpc_v1.Reply{
		Code:    0,
		Message: "OK",
	}, nil
}

/*
GetJob

Return JobInfo. If user wants the latest build, it will return the latest one. Otherwise it will return the latest 15 ones.
*/
func (s *server) GetJob(ctx context.Context, in *build_rpc_v1.Request) (*build_rpc_v1.JobInfo, error) {

	ji := new(build_rpc_v1.JobInfo)

	return ji, nil
}

/*
JobStatus

Update job status.
*/
func (s *server) JobStatus(ctx context.Context, in *build_rpc_v1.Builder) (*build_rpc_v1.Reply, error) {

	err := s.pg.UpdataBuildStatus(in.Status, in.Jid)
	if err != nil {
		return &build_rpc_v1.Reply{
			Code:    -1,
			Message: err.Error(),
		}, nil
	}

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

	build_rpc_v1.RegisterBuildServiceServer(s, &server{
		pg: store.PG(),
	})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
