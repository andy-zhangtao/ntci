package main

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	build_rpc_v1 "ntci/ci-grpc/build"
)

const (
	GitSuccess       = 1
	GitFailed        = -1
	NtciParseSuccess = 2
	NtciParseFailed  = -2
	Building         = 3
	BuildSuccess     = 4
	BuildFailed      = -4
)

//updateJobStatus
//Invoke build server for update job status.
//
//Job build has five flags:
//
// 1 - Git clone success
//-1 - Git clone failed
// 2 - Ntci parse success
//-2 - Ntci parse failed
// 3 - Building
// 4 - Build success
//-4 - Build failed
func updateJobStatus(flag int32) (err error) {
	conn, err := grpc.Dial(buildAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("Connect Build Service Error: %s", err.Error())
		return
	}

	defer conn.Close()

	c := build_rpc_v1.NewBuildServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.JobStatus(ctx, &build_rpc_v1.Builder{
		Jname:  gm.Name,
		Jid:    gm.Id,
		Status: flag,
		User:   gm.User,
	})

	if err != nil {
		logrus.Errorf("Invoke Build Service Error.  %v", err)
		return err
	}

	if r.Code != 0 {
		logrus.Errorf("Invoke Build Service Failed.  %d, %s", r.Code, r.Message)
		return errors.New("Invoke Build Service Failed ")
	}

	logrus.Debugf("Update Status Success: %d", r.Code)
	return
}
