package main

import (
	"fmt"
	"log"
	"encoding/json"
	  
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
	Seguimiento int64
	Tipo int64 //0:normal 1: prioritario 2: retail q
	Valor int64
	Intentos int64
	Estado int64 //0: En bodega 1: En Camino 2: Recibido 3: No Recibido
}

type Registro2 struct{
	IDpaquete string
	Seguimiento int64
	Tipo int64 //0:normal 1: prioritario 2: retail q
	Valor int64
	Intentos int64
	Estado int64//0: En bodega 1: En Camino 2: Recibido 3: No Recibido
}

var recibidos []Registro
var ganancia int64
var perdida int64
var total int64
func CalculoFinanza(){
	var ganancia =0
	var perdida =0
	var total =0
	for _,pack :=range recibidos{
		if pack.Estado==2{
			ganancia+=pack.Valor
		}else if pack.Estado==3{
			switch pack.Tipo{
			case 0:
			case 1:
				ganancia+=int64(float64(pack.Valor)*(0.3))
			case 2:
				ganancia+=pack.Valor
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
			recibidos=append(recibidos,m)
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever
}

func main(){

	RecepcionLogistica()
	CalculoFinanza()
	fmt.Println(recibidos)
	fmt.Println(ganancia,perdida,total)
}