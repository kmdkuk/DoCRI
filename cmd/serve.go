package cmd

import (
	"fmt"
	"github.com/kmdkuk/DoCRI/log"
	"github.com/kmdkuk/DoCRI/pkg/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"net"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "Serve a proxy to convert Docker REST API to CRI",
	Run: Serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Serve(_ *cobra.Command, _ []string) {
	fmt.Println("Hello, world")
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	dockerRuntimeServiceServer := &service.DockerRuntimeServiceServer{}
	runtimeapi.RegisterRuntimeServiceServer(server, dockerRuntimeServiceServer)
	err = server.Serve(listenPort)
	if err != nil {
		log.Fatal(err)
	}
}
