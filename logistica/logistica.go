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

type Registro struct{
	id_paquete string
	estado_paquete int64 //0: En bodega 1: En Camino 2: Recibido 3: No Recibido
	id_camion string
	seguimiento int64
	intentos int64
}

var numseg int64 =0
var fullreg [][] string
var colaretail[] Registro
var colaprioritario[] Registro
var colanormal[] Registro
func (s *server) Envio(ctx context.Context, msg *cl.EnvioRequest) (*cl.EnvioResponse, error){
	//fmt.Println(time.Now().Format("02-01-2006 15:04:05"),msg.GetId(), msg.GetProducto(), msg.GetValor(), msg.GetTienda(), msg.GetDestino(), msg.GetPrioritario(), numseg)
	var tipo string
	var seguimiento int64 = 0
	switch msg.GetPrioritario(){
	case 0:
		tipo="normal"
		seguimiento += numseg
		numseg+=1

	case 1:
		tipo="prioritario"
		seguimiento += numseg
		numseg+=1
	case 2:
		tipo="retail"

	default:
		break
	}
	
	linea:= []string{time.Now().Format("02-01-2006 15:04:05"),msg.GetId(),
			tipo,msg.GetProducto(),strconv.FormatInt(msg.GetValor(),10),msg.GetTienda(),msg.GetDestino(),
			strconv.FormatInt(seguimiento,10)} 
	reg := Registro{
		id_paquete : msg.GetId(),
		estado_paquete : 0,
		id_camion : "Sin Designar",
		seguimiento : seguimiento,
		intentos : 0,
	}
	switch msg.GetPrioritario(){
	case 0:
		colanormal= append(colanormal,reg)
	case 1:
		colaprioritario= append(colaprioritario,reg)
	case 2:
		colaretail= append(colaretail,reg)
	default:
		break
	}
	fullreg= append(fullreg,linea)
	//registro y sobreescritura por cada ingreso
	file,err:= os.OpenFile("registro.csv",os.O_CREATE|os.O_WRONLY,0777)
	defer file.Close()
	if err !=nil{
		os.Exit(1)
	}
	csvWriter:= csv.NewWriter(file)
	csvWriter.WriteAll(fullreg)
	csvWriter.Flush()
	fmt.Println(colaretail)
	fmt.Println(colaprioritario)
	fmt.Println(colanormal)
	return &cl.EnvioResponse {
		Msg: strconv.FormatInt(seguimiento,10),
	}, nil
}



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
	
}
