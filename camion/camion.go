package main
import (
	"context"
	"log"
	"fmt"
	"time"
	"reflect"
	pb "github.com/sirbernal/t1-sd2020-2/proto/camion_logistica"
	"google.golang.org/grpc"
	"math/rand"
)

const (
	dire = "localhost:50051"
)

type Envio struct{
	idPaquete string
	seguimiento int64
	tipo int64 //{normal, prioritario, retail}={0,1,2}
	valor int64
	intentos int64
	estado int64 //{En bodega, En camino, Recibido o No Recibido}={0,1,2,3}
	fecha_entrega time.Time
}

func menorEnvio(x Envio, y Envio)(Envio, Envio){
	if x.valor>y.valor{
		return x,y
	}else{
		return y,x
	}
}
func viaje(env [2]Envio)[2]Envio{
	void:= Envio{
		estado:1,
	}
	for i,pac := range env{
		if reflect.DeepEqual(pac,void){
			continue
		}
		for x:=0;x<4;x++{
			if x==3{
				pac.estado=3
				break
			}
			if pac.tipo<2 && int64(x)*10>pac.valor{
				fmt.Println("ta muy cara la weaita")
				pac.estado=3
				break
			}
			pac.intentos++
			rand.Seed(time.Now().UnixNano())
			probabilidad:= rand.Intn(101)
			if probabilidad<80{
				pac.estado=2
				pac.fecha_entrega=time.Now()
				break
			}else{
				fmt.Println("wea mala falló")
			}
		}
		env[i]=pac
	}
	return env
}
func simularEnvio(env [6]Envio)[6]Envio{ 
	e0,e1 := menorEnvio(env[0],env[1])
	e2,e3 := menorEnvio(env[2],env[3])
	e4,e5 := menorEnvio(env[4],env[5])
	cam1 :=[2]Envio{e0,e1} //camion retail 1
	cam2 :=[2]Envio{e2,e3}//camion retail 2
	cam3 :=[2]Envio{e4,e5} //camion normal
	cam1 = viaje(cam1)
	cam2 = viaje(cam2)
	cam3 = viaje(cam3)
	resultado := [6]Envio{cam1[0], cam1[1], cam2[0], cam2[1], cam3[0], cam3[1]} 
	return resultado

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
	var envios [6]Envio
	client := pb.NewCamionServiceClient(conn)
	stream, err := client.Camion(context.Background())
	waitc := make(chan struct{})

	msg := &pb.CamionRequest{IdPaquete : "",
		Seguimiento : -1,
		Tipo :0,
		Valor :  0,
		Intentos : 0,
		Estado : 0,}
	go func() {
		for i := 0; i < 1; i++ {
		stream.Send(msg) }
			
		/*for i := 1; i <= 10; i++ {
			log.Println("Sleeping...")
			time.Sleep(2 * time.Second)
			log.Println("Sending msg...")
			stream.Send(msg)
		}*/
	}()
	//time.Sleep(10 * time.Second)
	go func() {
		npack:= 0
		var resultado [6]Envio
		for {
			resp, err := stream.Recv()
			
			fmt.Println(resp.IdPaquete)
			fmt.Println(resp.Intentos)
			fmt.Println(npack)
			paquete := Envio{
				idPaquete : resp.IdPaquete,
				seguimiento : resp.Seguimiento,
				tipo :resp.Tipo,
				valor :  resp.Valor,
				intentos : resp.Intentos,
				estado : 1,
			}
			envios[npack]=paquete
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			npack++
			if npack>5{
				npack=0
				resultado = simularEnvio(envios)
				for _, pack := range resultado{
					msg2 := &pb.CamionRequest{
						IdPaquete: pack.idPaquete,   
						Seguimiento: pack.seguimiento,
						Tipo: pack.tipo,        
						Valor: pack.valor,       
						Intentos: pack.intentos,    
						Estado: pack.estado }
					stream.Send(msg2)
				}
			}
			time.Sleep(2 * time.Second)
			//fmt.Println(envios)
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

