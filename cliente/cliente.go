package main
import (
	"context"
	"fmt"
	"log"
	"time"
	cl "github.com/sirbernal/t1-sd2020-2/proto/cliente_logistica"
	"google.golang.org/grpc"
)

const (
	dire = "localhost:50051"
)

func main() {
	fmt.Println("Hola soy el usuario")
	conn, err := grpc.Dial(dire, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn err: %v", err)

	}
	defer conn.Close()
	c := cl.NewSeguimientoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	envio := &cl.SeguimientoRequest{
		Msg: "wea",
		Msg2: "wea2",
	}
	r, err := c.Seguimiento(ctx, envio)

	if err != nil {
		log.Fatalf("Requ err: %v", err)
	}

	log.Println("Respuesta : ", r.GetConfirmation())
}
