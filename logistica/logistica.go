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
	"encoding/json"
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
//por motivos de incompatibilidad de rabbit se genero esta segunda clase (más rapido que editar todo los valores)
type Registro2 struct{
	IDpaquete string
	Seguimiento int64
	Tipo int64 //0:normal 1: prioritario 2: retail q
	Valor int64
	Intentos int64
	Estado int64//0: En bodega 1: En Camino 2: Recibido 3: No Recibido
}


var numseg int64 =1 //contador de numero de seguimiento
var fullreg [][] string //lista que guarda todos registros de cliente
var colaretail[] Registro //cola que tiene los paquetes sin enviar de tipo retail
var colaprioritario[] Registro //cola que tiene los paquetes sin enviar de tipo prioritario
var colanormal[] Registro //cola que tiene los paquetes sin enviar de tipo normal
var completados[] Registro //lista que guarda todos los paquetes que pasaron por camion (entregados y no entregados)
var flag = false //variable utilizada para enviar paquete vacio a finanzas (notifica que el ultimo viaje de los camiones poseía al menos un paquete enviado)
func RemoveIndex(s []Registro, index int) []Registro { //funcion que elimina un item de una lista por su indice (tomada de internet)
	return append(s[:index], s[index+1:]...)
}

func CalcularEnvio() [6]Registro{ //genera una lista donde se enviaran los 6 paquetes para los 3 camiones correspondientes de la siguiente forma
	var paqtruck [6]Registro //{camionretail1,camionretail1,camionretail2,camionretail2,camionnormal,camionnormal}
	void := Registro{} //auxiliar para registros vacios
	for i,pack:= range colaretail{ //revisa primero la cola retail en intenta agregarlo en los primeros 4 cupos correspondientes
		if (!(reflect.DeepEqual(paqtruck[3],void))){ //si el cuarto espacio esta utilizado deja de agregar paquetes retail
			break
		}
		for a := 0; a < 4 ; a++{
			if reflect.DeepEqual(paqtruck[a],pack){ //si el paquete ya esta en cola no lo agrega
				break
			}
			if reflect.DeepEqual(paqtruck[a],void){ //si el espacio esta libre lo agrega
				pack.estado=1 //cambia el esta de "En bodega" a "En camion"
				paqtruck[a]=pack //guarda el paquete para enviarlo al camion
				colaretail[i]=pack //actualiza el estado de la cola
				break
			}
		}
	}
	for j,pack:= range colaprioritario{ //realiza lo mismo en los 6 espacios disponibles por ser tipo prioritario
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
	for k,pack:= range colanormal{ //realiza lo mismo solo en los dos espacios correspodiente para los tipo normal
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
func recepcionCamion(rescam [6]Registro){//funcion que actualiza colas al recibir el resultado de los camiones
	for _,pack := range rescam{//revisa cada paquete
		if reflect.DeepEqual(pack,Registro{estado:1})||reflect.DeepEqual(pack,Registro{}){ //si el paquete esta vacío lo omite
			continue
		}else{
			completados=append(completados,pack) //agrega el paquete a registro
			switch pack.tipo{ //busca en que cola esta el paquete para eliminarla de la misma
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
			PackageToFinanciero(pack) //envia el paquete a finanzas
			flag=true //notifica que al menos un paquete de los camiones existe
		}
	}
	if len(colanormal)+len(colaprioritario)+len(colaretail)==0 && len(completados)>0 &&flag{ //solo si las colas estan vacias, existen paquetes completados y el ultimo paquete fue real
		flag=false //reinicia condicional de paquetes reales
		PackageToFinanciero(Registro{}) //envia paquete vacio a finanzas lo cual notifica que ha terminado el sistema (ya envio todo)
		                    			//, donde puede realizar los calculos y actualización del archivo de finanzas
	}
}
func TranslateStatus(state int64)string{ //traductor a string del estado por numero
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
func BusquedaState(nseg int64)string{ //funcion que busca el estado de un envio
	if nseg==0{ //dado que retail tiene reservado este numero notifica lo mismo
		return "Seguimiento no disponible en Retail"
	}
	for _,pack:= range completados{ //primero revisa en la lista de entregados si se encuentra el envio
		if pack.seguimiento==nseg{
			return TranslateStatus(pack.estado)
		}
	}//revisa en el resto de colas como se encuentra el pedido
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
func BusquedaTruck(id string, valor string, seguimiento string)string{ //envia al camion el origen y destino del paquete que solicita (para seguir pauta y no alterar el paquete inicial)
	for _,j:= range fullreg{
		if id==j[1] && valor==j[4] && seguimiento==j[7]{
			return j[5]+","+j[6]
		}
	}
	return ""
}

func (s *server) Envio(ctx context.Context, msg *cl.EnvioRequest) (*cl.EnvioResponse, error){ //recepción de paquetes del cliente
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
	
	linea:= []string{time.Now().Format("02-01-2006 15:04:05"), //traduccion a lista de strings para escribir en archivo
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
	fullreg= append(fullreg,linea) //guarda la linea generada para el registro
	//registro y sobreescritura por cada ingreso
	file,err:= os.OpenFile("registro.csv",os.O_CREATE|os.O_WRONLY,0777) //actualiza archivo de registro
	defer file.Close()
	if err !=nil{
		os.Exit(1)
	}
	csvWriter:= csv.NewWriter(file)
	csvWriter.WriteAll(fullreg)
	csvWriter.Flush()
	CalcularEnvio() //actualiza la lista de espera
	return &cl.EnvioResponse {
		Msg: strconv.FormatInt(seguimiento,10),
	}, nil
	
	
}

func (s *server) Camion(stream pb.CamionService_CamionServer) error { //recibe la solicitud y envia 6 paquetes para los 3 camiones en espera
	recepcion:= [6]Registro{} //lista que guarda los paquetes que se enviaran a los camiones
	counterrecep:= 0 //contador de recepciones de paquetes de camiones
	fmt.Println(recepcion)
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Fatalf("RPC failed: %v", err)
		}
		if req.Seguimiento != -1{ //si los paquetes son reales procede a guardarlos en el registro de recepcion
			recepcion[counterrecep]=Registro{
				IDpaquete: req.IdPaquete,   
				seguimiento: req.Seguimiento,
				tipo: req.Tipo,        
				valor: req.Valor,       
				intentos: req.Intentos,    
				estado: req.Estado,	
			}
			counterrecep++
			if counterrecep==6{ //al concretar los 6 paquetes procede a actualizar las colas
				counterrecep=0
				recepcionCamion(recepcion) //funcion que actualiza las colas
				recepcion=[6]Registro{}
			}
		}else{//seguimiento=-1 representa un paquete vacío, es decir que camion esta disponible para hacer envios solicitando paquetes para enviar
			counterrecep=0 //reinicia el contador de paquetes
			paquetes := CalcularEnvio() //calcula que paquetes acorde a las colas iran a los camiones
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
				if err := stream.Send(&resp); err != nil { //envia los paquetes a camion
					log.Printf("send error %v", err)
				}
			}
		}
	}
}

func (s *server) Seguimiento(ctx context.Context, msg *cl.SeguimientoRequest) (*cl.SeguimientoResponse, error){ //funcion que recibe la solicitud de estado de seguimiento de cliente
	a, _ := strconv.Atoi(msg.GetSeguimiento())
	estado := BusquedaState(int64(a))
	return &cl.SeguimientoResponse {
		Estado: estado,
	}, nil
}

func (s *server) DatosCamion(ctx context.Context, msg *pb.DatosRequest) (*pb.DatosResponse, error){//funcion que recibe y envia datos extras a camion para registro (origen, destino)

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
func RegtoReg2(r Registro)Registro2{ //traductor de registro por compatibilidad con rabbit
	r2:= Registro2{
		IDpaquete : r.IDpaquete,
		Seguimiento : r.seguimiento,
		Tipo: r.tipo,
		Valor: r.valor,
		Intentos: r.intentos,
		Estado: r.estado,
	}
	return r2
}

func PackageToFinanciero(r Registro){//funcion que envia un registro de envio a finanzas
	
	conn, err := amqp.Dial("amqp://mqadmin:mqadminpassword@10.10.28.84:5672/") //conexion con finanzas por medio de rabbitmq
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	  
	//body := "Hello World!"

	var reg = RegtoReg2(r) //traduccion de registro por compatibilidad
	
	body, _ := json.Marshal(reg) //marshalling del registro a enviar
	

	err = ch.Publish( //envia el registro a finanzas
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "application/json",
		  Body:        []byte(body), //[]byte(body)
		})
	failOnError(err, "Failed to publish a message")

}


func main() {
	fmt.Println("Servidor PrestigioExpress <Logistica> corriendo")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error conectando: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCamionServiceServer(s, &server{}) //recibe conexión con el camion
	cl.RegisterEnvioServiceServer(s, &server{}) //recibe conexión con el cliente
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}

