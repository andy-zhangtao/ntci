package builder

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"ntci/ci-agent/dataBus"
	build_rpc_v1 "ntci/ci-grpc/build"
)

func InvokeBuilderServiceRun(req *build_rpc_v1.Request) (res *build_rpc_v1.Reply, err error) {
	bus := dataBus.GetBus()

	conn, err := grpc.Dial(bus.Build[bus.BuildMode].Addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("did not connect: %v", err))
	}

	defer conn.Close()

	c := build_rpc_v1.NewBuildServiceClient(conn)

	return c.Run(context.Background(), req)

}

func InvokeBuilderServiceRestart(req *build_rpc_v1.Request) (res *build_rpc_v1.Reply, err error) {
	bus := dataBus.GetBus()

	conn, err := grpc.Dial(bus.Build[bus.BuildMode].Addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("did not connect: %v", err))
	}

	defer conn.Close()

	c := build_rpc_v1.NewBuildServiceClient(conn)

	return c.RestartJob(context.Background(), req)
}
