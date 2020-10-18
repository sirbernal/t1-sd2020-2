package main

import (
	"fmt"
	"log"
	  
    "github.com/streadway/amqp"
/*	"context"
	"log"
	"net"
	
	"time"
	"strconv"
	"os"
	"encoding/csv"
	"reflect"
	cl "github.com/sirbernal/t1-sd2020-2/proto/cliente_logistica"
	pb "github.com/sirbernal/t1-sd2020-2/proto/camion_logistica"
	grpc "google.golang.org/grpc"*/
)

type Registro struct{
	IDpaquete string
	seguimiento int64
	tipo int64 //0:normal 1: prioritario 2: retail q
	valor int64
	intentos int64
	estado int64 //0: En bodega 1: En Camino 2: Recibido 3: No Recibido
}

var recibidos []Registro
var ganancia int64
var perdida int64
var total int64
func CalculoFinanza(){
	for _,pack :=range recibidos{
		if pack.estado==2{
			ganancia+=pack.valor
		}else if pack.estado==3{
			switch pack.tipo{
			case 0:
			case 1:
				ganancia+=int64(float64(pack.valor)*(0.3))
			case 2:
				ganancia+=pack.valor
			default:
				continue
			}
		}
		if pack.intentos>1{
			perdida+=(10*(pack.intentos-1))
		}
		fmt.Println(ganancia,perdida,total)
	}
	
	total=ganancia-perdida
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func PrintHolamundo(){
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
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever


}

func main(){

	PrintHolamundo()
	recibidos=append(recibidos,Registro{"",1,0,10,1,2},
		Registro{"",2,1,30,3,3},Registro{"",3,2,40,2,2})
	CalculoFinanza()
	fmt.Println(recibidos)
	fmt.Println(ganancia,perdida,total)
}