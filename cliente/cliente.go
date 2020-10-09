package main
import (
	"context"
	"fmt"
	"log"

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

	req := &cl.SeguimientoRequest {
		Msg : "Mi tula",
    	Msg2 : "Mi ano",
	}
	
	res, err := c.Seguimiento(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call Sum function: %v", err)
	}

	log.Printf("Respuesta: %v", res.GetConfirmation())
}
