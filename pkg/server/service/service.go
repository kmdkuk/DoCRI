package service

import (
	"context"
	"log"

	"github.com/docker/docker/client"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

type DockerRuntimeServiceServer struct {
	runtimeapi.UnimplementedRuntimeServiceServer
}

func (c *DockerRuntimeServiceServer) Version(context.Context, *runtimeapi.VersionRequest) (*runtimeapi.VersionResponse, error) {
	log.Println("リクエストがきたぞえ")
	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	version, err := cli.ServerVersion(ctx)
	if err != nil {
		panic(err)
	}
	log.Println(version)
	return &runtimeapi.VersionResponse{}, nil
}
