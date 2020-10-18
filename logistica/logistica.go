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
	"reflect"
	"strings"
	cl "github.com/sirbernal/t1-sd2020-2/proto/cliente_logistica"
	pb "github.com/sirbernal/t1-sd2020-2/proto/camion_logistica"
	grpc "google.golang.org/grpc"
	"github.com/streadway/amqp"
)

const (
	port = ":50051"
)



type server struct {
}

type Registro struct{
	IDpaquete string
	seguimiento int64
	tipo int64 //0:normal 1: prioritario 2: retail q
	valor int64
	intentos int64
	estado int64 //0: En bodega 1: En Camino 2: Recibido 3: No Recibido
}


var numseg int64 =1
var fullreg [][] string
var colaretail[] Registro
var colaprioritario[] Registro
var colanormal[] Registro
var completados[] Registro

func RemoveIndex(s []Registro, index int) []Registro {
	return append(s[:index], s[index+1:]...)
}
func CalcularEnvio() [6]Registro{
	var paqtruck [6]Registro
	void := Registro{}
	for i,pack:= range colaretail{
		if (!(reflect.DeepEqual(paqtruck[3],void))){
			break
		}
		for a := 0; a < 4 ; a++{
			if reflect.DeepEqual(paqtruck[a],pack){
				break
			}
			if reflect.DeepEqual(paqtruck[a],void){
				pack.estado=1
				paqtruck[a]=pack
				colaretail[i]=pack
				break
			}
		}
	}
	for j,pack:= range colaprioritario{
		if (!(reflect.DeepEqual(paqtruck[5],void))){
				break
		}
		for b := 0 ;b<6;b++{
			if reflect.DeepEqual(paqtruck[b],pack){
				break
			}
			if reflect.DeepEqual(paqtruck[b],void){
				pack.estado=1
				paqtruck[b]=pack
				colaprioritario[j]=pack
				break
			}
		}
	}
	for k,pack:= range colanormal{
		if (!(reflect.DeepEqual(paqtruck[5],void))){
				break
		}
		for c := 4; c < 6; c++{
			if reflect.DeepEqual(paqtruck[c],pack){
				break
			}
			if reflect.DeepEqual(paqtruck[c],void){
				pack.estado=1
				paqtruck[c]=pack
				colanormal[k]=pack
				break
			}
		}
	}
	return paqtruck
}
func recepcionCamion(rescam [6]Registro){
	for _,pack := range rescam{
		if reflect.DeepEqual(pack,Registro{estado:1}){
			continue
		}else{
			completados=append(completados,pack)
			switch pack.tipo{
			case 0:
				for i,j :=range colanormal{
					if j.seguimiento==pack.seguimiento && j.IDpaquete==pack.IDpaquete && j.valor==pack.valor{
						colanormal=RemoveIndex(colanormal,i)
						break
					}
				}
			case 1:
				for i,j :=range colaprioritario{
					if j.seguimiento==pack.seguimiento && j.IDpaquete==pack.IDpaquete && j.valor==pack.valor{
						colaprioritario=RemoveIndex(colaprioritario,i)
						break
					}
				}
			case 2:
				for i,j :=range colaretail{
					if j.seguimiento==pack.seguimiento && j.IDpaquete==pack.IDpaquete && j.valor==pack.valor{
						colaretail=RemoveIndex(colaretail,i)
						break
					}
				}
			default:
				continue
			}
		}
	}
}
func TranslateStatus(state int64)string{
	switch state{
	case 0:
		return "En bodega"
	case 1:
		return "En Camino"
	case 2:
		return "Recibido"
	case 3:
		return "No Recibido"
	default:
		return ""
	}
}
func BusquedaState(nseg int64)string{
	if nseg==0{
		return "Seguimiento no disponible en Retail"
	}
	for _,pack:= range completados{
		if pack.seguimiento==nseg{
			return TranslateStatus(pack.estado)
		}
	}
	for _,pack:= range colaprioritario{
		if pack.seguimiento==nseg{
			return TranslateStatus(pack.estado)
		}
	}
	for _,pack:= range colanormal{
		if pack.seguimiento==nseg{
			return TranslateStatus(pack.estado)
		}
	}
	return "Número de seguimiento inexistente en sistema"
}
func BusquedaTruck(id string, valor string, seguimiento string)string{
	for _,j:= range fullreg{
		if id==j[1] && valor==j[4] && seguimiento==j[7]{
			return j[5]+","+j[6]
		}
	}
	return ""
}

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
	
	linea:= []string{time.Now().Format("02-01-2006 15:04:05"),
			msg.GetId(),tipo,msg.GetProducto(),
			strconv.FormatInt(msg.GetValor(),10),
			msg.GetTienda(),msg.GetDestino(),
			strconv.FormatInt(seguimiento,10)} 
	reg := Registro{
		IDpaquete : msg.GetId(),
		seguimiento : seguimiento,
		tipo : msg.GetPrioritario(),
		valor :  msg.GetValor(),
		intentos : 0,
		estado : 0,
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
	//fmt.Println(colaretail)
	//fmt.Println(colaprioritario)
	//fmt.Println(colanormal)
	fmt.Println(CalcularEnvio())
	return &cl.EnvioResponse {
		Msg: strconv.FormatInt(seguimiento,10),
	}, nil
	
	
}

func (s *server) Camion(stream pb.CamionService_CamionServer) error {
	recepcion:= [6]Registro{}
	counterrecep:= 0
	fmt.Println(recepcion)
	for {
		//fmt.Println(BusquedaState(10))
		//fmt.Println(BusquedaTruck("SA7748JK","Cama","115","0"))
		req, err := stream.Recv()
		if err != nil {
			log.Fatalf("RPC failed: %v", err)
		}
		if req.Seguimiento != -1{
			//fmt.Println("Soy un paquete de verdad")
			//fmt.Println(counterrecep)
			recepcion[counterrecep]=Registro{
				IDpaquete: req.IdPaquete,   
				seguimiento: req.Seguimiento,
				tipo: req.Tipo,        
				valor: req.Valor,       
				intentos: req.Intentos,    
				estado: req.Estado,	
			}
			/*fmt.Println(Registro{
				IDpaquete: req.IdPaquete,   
				seguimiento: req.Seguimiento,
				tipo: req.Tipo,        
				valor: req.Valor,       
				intentos: req.Intentos,    
				estado: req.Estado,	
			})*/
			counterrecep++
			if counterrecep==6{
				counterrecep=0
				fmt.Println(recepcion)
				recepcionCamion(recepcion)
				recepcion=[6]Registro{}
				fmt.Println(len(colaretail))
				fmt.Println(len(colaprioritario))
				fmt.Println(len(colanormal))
				fmt.Println(len(completados))

			}
		}else{
			//fmt.Println("Paquetito fake")
			counterrecep=0
			paquetes := CalcularEnvio()
			for _, paquete:= range paquetes {
				//fmt.Println(paquete)
				resp := pb.CamionResponse{
					IdPaquete: paquete.IDpaquete,   
					Seguimiento: paquete.seguimiento,
					Tipo: paquete.tipo,        
					Valor: paquete.valor,       
					Intentos: paquete.intentos,    
					Estado: paquete.estado,	
				}
				if err := stream.Send(&resp); err != nil {
					log.Printf("send error %v", err)
				}
			}
		}
	}
}

func (s *server) Seguimiento(ctx context.Context, msg *cl.SeguimientoRequest) (*cl.SeguimientoResponse, error){
	a, _ := strconv.Atoi(msg.GetSeguimiento())
	estado := BusquedaState(int64(a))
	return &cl.SeguimientoResponse {
		Estado: estado,
	}, nil
}

func (s *server) DatosCamion(ctx context.Context, msg *pb.DatosRequest) (*pb.DatosResponse, error){

	name := msg.GetName()
	n := strings.Split(name,",")
	
	return &pb.DatosResponse {
		Dato: BusquedaTruck(n[0],n[1],n[2]),
	}, nil

}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
  }

func Holamundo(){
	
	conn, err := amqp.Dial("amqp://mqadmin:mqadminpassword@10.10.28.84:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	  
	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "text/plain",
		  Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

}


func main() {
	fmt.Println("Servidor PrestigioExpress <Logistica> corriendo")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error conectando: %v", err)
	}

	s := grpc.NewServer() */
	Holamundo()
	cl.RegisterSeguimientoServiceServer(s, &server{})
	pb.RegisterCamionServiceServer(s, &server{})
	cl.RegisterEnvioServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}

