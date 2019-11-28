package main

import (
	"context"
	"errors"
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

func (s *server) RestartJob(ctx context.Context, in *build_rpc_v1.Request) (*build_rpc_v1.Reply, error) {
	logrus.Debugf("Receive Restart Request. User: %s Name: %s Branch: %s Git: %s ID: %d Language: %s Ver: %s. Sha: %s Message: %s ", in.User, in.Name, in.Branch, in.Url, in.Id, in.Language, in.Lanversion, in.Sha, in.Message)

	b := store.Build{
		Name:      in.Name,
		Branch:    in.Branch,
		Git:       in.Url,
		Timestamp: time.Now().Local(),
		Token:     bus.Token,
		Addr:      bus.Addr,
		User:      in.User,
		Sha:       in.Sha,
		Message:   in.Message,
		Id:        int(in.Id),
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

	err := deploy.DeleteJob(b)
	if err != nil {
		logrus.Errorf("Delete Job Error: %s", err.Error())
		return &build_rpc_v1.Reply{
			Code:    -1,
			Message: fmt.Sprintf("Delete Job Error: %s", err.Error()),
		}, nil
	}

	err = deploy.NewJob(b, in.Env)
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

	logrus.Debugf("Receive Build Request. User: %s Name: %s Branch: %s Git: %s ID: %d Language: %s Ver: %s. Sha: %s Message: %s ", in.User, in.Name, in.Branch, in.Url, in.Id, in.Language, in.Lanversion, in.Sha, in.Message)

	b := store.Build{
		Name:       in.Name,
		Branch:     in.Branch,
		Git:        in.Url,
		Timestamp:  time.Now().Local(),
		Token:      bus.Token,
		Addr:       bus.Addr,
		User:       in.User,
		Sha:        in.Sha,
		Message:    in.Message,
		Id:         int(in.Id),
		Dockerfile: in.Dockerfile,
		BuildScript: in.BuildScript,
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

	err := deploy.NewJob(b, in.Env)
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
func (s *server) GetJob(ctx context.Context, in *build_rpc_v1.JobRequest) (*build_rpc_v1.JobInfo, error) {

	var jd []*build_rpc_v1.JobDetail

	bs, err := s.pg.GetBuild(in.Owner, in.Name)
	if err != nil {
		return nil, err
	}

	for _, b := range bs {
		jd = append(jd, &build_rpc_v1.JobDetail{
			Name:      b.Name,
			Status:    int32(b.Status),
			Timestamp: b.Timestamp.Format("2006-01-02 15:04:05"),
			Branch:    b.Branch,
			Url:       b.Git,
			Id:        int32(b.Id),
			Sha:       b.Sha,
			Message:   b.Message,
		})
	}

	return &build_rpc_v1.JobInfo{
		Count: int32(len(jd)),
		Jd:    jd,
	}, nil
}

/*
JobStatus

Update job status.
*/
func (s *server) JobStatus(ctx context.Context, in *build_rpc_v1.Builder) (*build_rpc_v1.Reply, error) {

	err := s.pg.UpdataBuildStatus(in.Status, in.Jname, in.Jid, in.User)
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

func (s *server) GetJobLog(in *build_rpc_v1.Job, ls build_rpc_v1.BuildService_GetJobLogServer) (err error) {

	err = deploy.GetJobLog(in.Name, true, ls)
	if err != nil && err.Error() == "EOF" {
		return errors.New("EOF")
	}

	if err != nil && err.Error() != "EOF" {
		logrus.Error(err)
		return
	}

	return
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
