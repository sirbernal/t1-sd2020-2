package main
import (
	"context"
	"log"
	"time"
	pb "github.com/sirbernal/t1-sd2020-2/proto/camion_logistica"
	"google.golang.org/grpc"
)

const (
	dire = "localhost:50051"
)

type Envio struct{
	id string
	producto string
	valor int64
	tienda string
	destino string
	prioritario int64 //{normal, prioritario, retail}={0,1,2}
}

func main() {
	
	//fmt.Println("Usuario del sistema PrestigioExpress, por favor ingresar el tiempo de envio de pedidos a logistica (en segundos): ")
	//var tiempo int64
	//fmt.Scanln(&tiempo)  
	// Creando conexion TCP
	conn, err := grpc.Dial(dire, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn err: %v", err)
	}
	
	client := pb.NewCamionServiceClient(conn)
	stream, err := client.Camion(context.Background())
	waitc := make(chan struct{})

	msg := &pb.CamionRequest{Status: "wea"}
	go func() {
		for i := 1; i <= 10; i++ {
			log.Println("Sleeping...")
			time.Sleep(2 * time.Second)
			log.Println("Sending msg...")
			stream.Send(msg)
		}
	}()
	//time.Sleep(10 * time.Second)
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			log.Printf(resp.Algo)
		}
	}()


	<-waitc
	stream.CloseSend()
}


	/* c := cl.NewSeguimientoServiceClient(conn)
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

	log.Println("Respuesta : ", r.GetConfirmation()) */

