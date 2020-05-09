package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kmdkuk/DoCRI/pkg/server/service"
	"google.golang.org/grpc"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func main() {
	fmt.Println("Hello, world")
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	dockerRuntimeServiceServer := &service.DockerRuntimeServiceServer{}
	runtimeapi.RegisterRuntimeServiceServer(server, dockerRuntimeServiceServer)
	server.Serve(listenPort)
}
