package main

import (
	"context"
	"log"
	"net"
	"fmt"
	cl "github.com/sirbernal/t1-sd2020-2/proto/cliente_logistica"
	grpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Pedido struct{
	id string
	producto string
	valor int32
	tienda string
	destino string
	prioritario int32
}

type server struct {
}

/* func (s *server) Seguimiento(ctx context.Context, msg *cl.SeguimientoRequest) (*cl.SeguimientoResponse, error){
	fmt.Println(msg.GetMsg(), msg.GetMsg2())
	return &cl.SeguimientoResponse{ Confirmation: "Chupa el pico",}, nil
} */

func (s *server) Envio(ctx context.Context, msg *cl.EnvioRequest) (*cl.EnvioResponse, error){
	fmt.Println(msg.GetId(), msg.GetProducto(), msg.GetValor(), msg.GetTienda(), msg.GetDestino())
	return &cl.EnvioResponse {
		Msg: "Recibido!",
	}, nil
}

/*
func Seguimiento(ctx context.Context, in *cl.SeguimientoRequest) (*cl.SeguimientoResponse, error) {
	fmt.Println("Wea llego a server")
	return &cl.SeguimientoResponse{}, nil
}
*/


func main() {
	fmt.Println("Servidor PrestigioExpress <Logistica> corriendo")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error conectando: %v", err)
	}


	s := grpc.NewServer()

	//cl.RegisterSeguimientoServiceServer(s, &server{})
	cl.RegisterEnvioServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("Wea lista pa hacer algo")
}
