package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"time"
	"strconv"
	"os"
	"encoding/csv"
	cl "github.com/sirbernal/t1-sd2020-2/proto/cliente_logistica"
	grpc "google.golang.org/grpc"
)

const (
	port = ":50051"
)



type server struct {
}

var numseg int64 =0
var fullreg [][] string


func (s *server) Envio(ctx context.Context, msg *cl.EnvioRequest) (*cl.EnvioResponse, error){
	//fmt.Println(time.Now().Format("02-01-2006 15:04:05"),msg.GetId(), msg.GetProducto(), msg.GetValor(), msg.GetTienda(), msg.GetDestino(), msg.GetPrioritario(), numseg)
	var tipo string
	switch msg.GetPrioritario(){
	case 0:
		tipo="normal"
	case 1:
		tipo="prioritario"
	case 2:
		tipo="retail"
	default:
		break
	}
	linea:= []string{time.Now().Format("02-01-2006 15:04:05"),msg.GetId(),
			tipo,msg.GetProducto(),strconv.FormatInt(msg.GetValor(),10),msg.GetTienda(),msg.GetDestino(),
			strconv.FormatInt(numseg,10)} 
	fullreg= append(fullreg,linea)
	numseg+=1
	//registro y sobreescritura por cada ingreso
	file,err:= os.OpenFile("registro.csv",os.O_CREATE|os.O_WRONLY,0777)
	defer file.Close()
	if err !=nil{
		os.Exit(1)
	}
	csvWriter:= csv.NewWriter(file)
	csvWriter.WriteAll(fullreg)
	csvWriter.Flush()
	
	return &cl.EnvioResponse {
		Msg: strconv.FormatInt(numseg-1,10),
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
