package main

import (
	"context"
	"log"
	"net"

	"github.com/nikkefimov/grpc/go_service"
	"google.golang.org/grpc"
)

type go_serviceServer struct {
	go_service.UnimplementedGoServiceServer
}

func (s go_serviceServer) Create(context.Context, *go_service.CreateRequest) (*go_service.CreateResponse, error) {
	return &go_service.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("can't create listner: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &go_serviceServer{}

	go_service.RegisterGoServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("serve error: %s", err)
	}

}
