package main

import (
	"context"

	"github.com/nikkefimov/grpc/go_service"
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

}
