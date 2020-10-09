package main

import (
	"context"
	"fmt"
	"log"
	"net"
	protos "t1-sd2020-2/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
}

func Seguimiento(ctx context.Context, in *SeguimientoRequest) (*SeguimientoResponse, error) {
	fmt.Println("Wea llego a server")
	return &SeguimientoResponse{}, nil
}

func main() {
	lis, err := net.Listen("tpc", port)
	if err != nil {
		log.Fatal("erro %v", err)
	}

	defer conn.Close()

	s := grpc.NewServer()

	protos.RegisterSeguimientoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
