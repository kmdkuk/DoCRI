package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("client connection error:", err)
	}
	defer conn.Close()
	client := runtimeapi.NewRuntimeServiceClient(conn)
	message := &runtimeapi.VersionRequest{}
	res, err := client.Version(context.TODO(), message)
	log.Printf("result:%#v \n", res)
	log.Printf("error::%#v \n", err)
}
