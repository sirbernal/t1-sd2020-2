package main

import (
	"fmt"
	"log"
	"encoding/json"
	"reflect"
	"github.com/streadway/amqp"
	"encoding/csv"
	"os"
	"strconv"
/*	"context"
	"log"
	"net"
	
	"time"
	"encoding/csv"
	
	cl "github.com/sirbernal/t1-sd2020-2/proto/cliente_logistica"
	pb "github.com/sirbernal/t1-sd2020-2/proto/camion_logistica"
	grpc "google.golang.org/grpc"*/
)

type Registro struct{
	IDpaquete string
	Seguimiento int64
	Tipo int64 //0:normal 1: prioritario 2: retail q
	Valor int64
	Intentos int64
	Estado int64 //0: En bodega 1: En Camino 2: Recibido 3: No Recibido
}



var recibidos []Registro
var resumen [][]string
var ganancia int64
var perdida int64
var total int64
func EscribirArchivoFinanza(){
	archivo:="resumenfinanzas.csv"
	file,err:= os.OpenFile(archivo,os.O_CREATE|os.O_WRONLY,0777)
	defer file.Close()
	if err !=nil{
		os.Exit(1)
	}
	csvWriter:= csv.NewWriter(file)
	csvWriter.WriteAll(resumen)
	csvWriter.Flush()
	
}
func CalculoFinanza(){
	var resumentemp [][]string
	resumentemp=append(resumentemp,[]string{"Id Paquete","Estado","Intentos","Ganancia","Perdida"})
	ganancia =0
	perdida =0
	total =0
	for _,pack :=range recibidos{
		var gananciapack int64
		var perdidapack int64
		var state string
		if pack.Estado==2{
			ganancia+=pack.Valor
			gananciapack+=pack.Valor
			state="Envío Completado"
		}else if pack.Estado==3{
			state="Envío no Entregado"
			switch pack.Tipo{
			case 0:
			case 1:
				ganancia+=int64(float64(pack.Valor)*(0.3))
				gananciapack+=int64(float64(pack.Valor)*(0.3))
			case 2:
				ganancia+=pack.Valor
				gananciapack+=pack.Valor
			default:
				continue
			}
		}
		if pack.Intentos>1{
			perdida+=(10*(pack.Intentos-1))
			perdidapack+=(10*(pack.Intentos-1))
		}
		linea:=[]string{pack.IDpaquete,state,strconv.FormatInt(pack.Intentos,10),strconv.FormatInt(gananciapack,10),strconv.FormatInt(perdidapack,10)}
		resumentemp=append(resumentemp,linea)
		//fmt.Println(ganancia,perdida,total)
	}
	resumen=resumentemp
	total=ganancia-perdida
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func RecepcionLogistica(){
	conn, err := amqp.Dial("amqp://mqadmin:mqadminpassword@localhost:5672/")
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


	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")
	  
	forever := make(chan bool)
	  
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var m Registro

			_ = json.Unmarshal(d.Body, &m)

			fmt.Println(m)
			if reflect.DeepEqual(m,Registro{}){
				CalculoFinanza()
				EscribirArchivoFinanza()
				fmt.Println(ganancia,perdida,total)
			}
			recibidos=append(recibidos,m)
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever
}

func main(){

	RecepcionLogistica()
}