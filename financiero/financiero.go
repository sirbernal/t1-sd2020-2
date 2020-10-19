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
)

type Registro struct{
	IDpaquete string
	Seguimiento int64
	Tipo int64 //0:normal 1: prioritario 2: retail q
	Valor int64
	Intentos int64
	Estado int64 //0: En bodega 1: En Camino 2: Recibido 3: No Recibido
}



var recibidos []Registro //guardado de todos los registros recibidos de la cola de rabbit
var resumen [][]string //arreglo que guarda las lineas del resumen solicitado a documentar
//variables solicitadas en enunciado
var ganancia int64
var perdida int64
var total int64
func EscribirArchivoFinanza(){ //escribe el archivo solicitado que posee por linea el detalle de cada envio
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
func CalculoFinanza(){ //realiza los calculos de cada paquete recibido en finanzas
	var resumentemp [][]string //auxiliar de reinicio
	resumentemp=append(resumentemp,[]string{"Id Paquete","Estado","Intentos","Ganancia","Perdida"}) //primera linea para encabezado del documento
	//reinicio de variables solicitadas (calculara todo desde el comienzo)
	ganancia =0
	perdida =0
	total =0
	for _,pack :=range recibidos{
		//variables para calculo de ganancia y perdida del envio en especifico
		var gananciapack int64
		var perdidapack int64
		var state string
		if pack.Estado==2{ //si el envio se concreta
			ganancia+=pack.Valor
			gananciapack+=pack.Valor
			state="Envío Completado"
		}else if pack.Estado==3{ //si el envio no se concreta
			state="Envío no Entregado"
			switch pack.Tipo{ //verifica por tipo las ganancias que calculará
			case 0: //caso normal
			case 1: //caso prioritario
				ganancia+=int64(float64(pack.Valor)*(0.3))
				gananciapack+=int64(float64(pack.Valor)*(0.3))
			case 2: //caso retail
				ganancia+=pack.Valor
				gananciapack+=pack.Valor
			default: 
				continue
			}
		}
		if pack.Intentos>1{ //calculo de perdidas por cantidad de envios
			perdida+=(10*(pack.Intentos-1))
			perdidapack+=(10*(pack.Intentos-1))
		}
		linea:=[]string{pack.IDpaquete,state,strconv.FormatInt(pack.Intentos,10),strconv.FormatInt(gananciapack,10),strconv.FormatInt(perdidapack,10)} //genera linea para documento
		resumentemp=append(resumentemp,linea) //agrega la linea al resumen
		//fmt.Println(ganancia,perdida,total)
	}
	//actualiza el resumen del documento y el total calculado
	resumen=resumentemp
	total=ganancia-perdida
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func RecepcionLogistica(){//funcion encargada de recibir los paquetes de la cola de rabbit desde logistica
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
		for d := range msgs { //revisa cada mensaje recibido en cola
			log.Printf("Received a message: %s", d.Body)

			var m Registro

			_ = json.Unmarshal(d.Body, &m) //unmarchall del mensaje

			if reflect.DeepEqual(m,Registro{}){ //si el mensaje es vacio realiza lo siguiente (especificado en logistica los casos posibles)
				CalculoFinanza() //realiza los calculos de todos los paquetes
				EscribirArchivoFinanza() //actualiza el archivo de finanzas
				fmt.Println(ganancia,perdida,total) //imprime las ganancias, perdidas y total
			}else{
				recibidos=append(recibidos,m) //guarda el mensaje
			}
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever
}

func main(){
	RecepcionLogistica()
}