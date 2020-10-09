package main
import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	dire = "localhost:50051"
)

func main() {
	fms.Println("Hola soy el usuario")
	conn, err := grpc.Dial(dire, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn err: %v", err)

	}
	defer conn.Close()
	c := NewSeguimientoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	envios := []*Envio{}

	envios = append(products, &Product){
		msg : "wea",
		msg2: "wea2"
	}
	r, err := c.SeguimientoService(ctx, &SeguimientoServiceRequest{
		
		Envios : envios
	}

	if err =! nil {
		log.Fatalf("Requ err: %v", err)
	}

	log.Println("Respuesta : ", r)
}
