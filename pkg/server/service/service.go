package service

import (
	"context"
	"log"

	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

type DockerRuntimeServiceServer struct {
	runtimeapi.UnimplementedRuntimeServiceServer
}

func (c *DockerRuntimeServiceServer) Version(context.Context, *runtimeapi.VersionRequest) (*runtimeapi.VersionResponse, error) {
	log.Println("リクエストがきたぞえ")
	return &runtimeapi.VersionResponse{}, nil
}
