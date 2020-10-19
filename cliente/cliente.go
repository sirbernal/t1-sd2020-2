package main
import (
	"context"
	"fmt"
	"log"
	"time"
	"encoding/csv"
	"os"
	"strconv"
	"io"

	cl "github.com/sirbernal/t1-sd2020-2/proto/cliente_logistica"
	"google.golang.org/grpc"
)

const (
	//10.10.28.82
	dire = "10.10.28.82:50051"
)

type Envio struct{
	id string
	producto string
	valor int64
	tienda string
	destino string
	prioritario int64 //{normal, prioritario, retail}={0,1,2}
}

var tiempopaquetes string //tiempo de espera+auxiliar de conversion
var tiempopaquetesint int
func envioRegistro(archivo string)(){ //funcion que separa el archivo linea por linea para ser enviado a logistica
	conn, err := grpc.Dial(dire, grpc.WithInsecure()) //inicia conexión con logistica 
	if err != nil {
		log.Fatalf("Conn err: %v", err)

	}
	defer conn.Close()
	doc :=archivo
	
	f, err := os.Open(doc) 
	if err != nil {
		log.Fatalf("Error al abrir el CSV: %v", err)
	}
	
	reader := csv.NewReader(f)
	reader.Comma = ','         
	reader.FieldsPerRecord = 6
	var lista_envios []Envio //guarda en formato de Envio cada linea

	for { //recorre archivo linea por linea para enviarlo
		record, err := reader.Read()

		if err == io.EOF { // cuando termina de leer el archivo
			break 
		}
		
		if record[0] == "id"{ //saltar primera linea
			continue
		}

		v, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Error al procesar dato: %v", err)
			continue
		}
		var p int64 =2
		if len(record)==6{
			prior, err := strconv.ParseInt(record[5], 10, 64)
				if err != nil {
						log.Fatalf("Error al procesar dato: %v", err)
						continue
				}
			p=prior
		}
		//fmt.Println(p)
		envio := Envio{
			id : record[0],
			producto : record[1],
			valor : v,
			tienda : record[3],
			destino : record[4],
			prioritario: p,
		}
		lista_envios = append(lista_envios, envio)//guarda la nueva linea
	}

	//fmt.Println(lista_envios)
	//Envio de retail
	
	for i := 0; i < len(lista_envios); i++ { //recorre las lineas guardadas en el nuevo formato
		
		c := cl.NewEnvioServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		envio := &cl.EnvioRequest{
			Id : lista_envios[i].id,
			Producto : lista_envios[i].producto,
			Valor : lista_envios[i].valor,
			Tienda : lista_envios[i].tienda,
			Destino : lista_envios[i].destino,
			Prioritario : lista_envios[i].prioritario,
		}
		
		r, err := c.Envio(ctx, envio) //envia la linea a logistica
		time.Sleep(time.Duration(tiempopaquetesint*1000)*time.Millisecond)
		if err != nil {
			log.Fatalf("Requ err: %v", err)
		}
		
		if (r.GetMsg()!="0"){
			fmt.Println("Codigo de seguimiento: "+r.GetMsg()) //Imprime el código de seguimiento solo si no es de tipo retail
		}
	}	
}

func ShowSeguimiento (){ //solicita el estado de un paquete en base al numero de seguimiento a logistica
	conn, err := grpc.Dial(dire, grpc.WithInsecure()) //inicia conexion con logistica
	if err != nil {
		log.Fatalf("Conn err: %v", err)

	}
	defer conn.Close()
	fmt.Println(" Ingrese codigo de seguimiento:  ")
	var seguimiento string
	fmt.Scanln(&seguimiento)
	c := cl.NewEnvioServiceClient(conn) //envia número de seguimiento
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	seg := &cl.SeguimientoRequest{
		Seguimiento: seguimiento,
	}

	r, _ := c.Seguimiento(ctx, seg) //recibe estado del pedido

	fmt.Println("Estado del pedido: ", r.GetEstado())
	fmt.Println("")
}

func main() {
	fmt.Print("Inicializando servicio cliente...\nIngrese tiempo (segundos) de envio de paquetes:")
	for{
		fmt.Scanln(&tiempopaquetes)
		intmode, err:=strconv.Atoi(tiempopaquetes)
		if err!=nil{
			fmt.Print("Formato no válido, ingrese nuevamente:")
			continue
		}
		tiempopaquetesint=intmode
		break
	}
	var menu string
	
	for {
		fmt.Print("Menu Cliente \n1.-Modo retail\n2.-Modo Pyme\n3.-Realizar Seguimiento\nIngrese opción:")
		_,err:=fmt.Scanln(&menu)
		if err!=nil{
			fmt.Print("\nFormato de ingreso no válido, pruebe nuevamente:")
			continue
		}
		switch menu{
		case "1":
			for{
				fmt.Print("Menu Retail \n1.-Enviar registros\n2.-Volver atrás\nIngrese opción:")
				_,err:=fmt.Scanln(&menu)
				if err!=nil{
					fmt.Print("\nFormato de ingreso no válido, pruebe nuevamente:")
					continue
				}
				switch menu{
				case "1":
					envioRegistro("retail.csv")
				case "2":
					break
				default:
					fmt.Print("\nFormato u opción no válida, pruebe nuevamente:\n\n")
					continue
				}	
			}
		case "2":
			for{
				fmt.Print("Menu Pymes \n1.-Enviar registros\n2.-Realizar Seguimiento\n3.-Volver\nIngrese opción:")
				_,err:=fmt.Scanln(&menu)
				if err!=nil{
					fmt.Print("\nFormato de ingreso no válido, pruebe nuevamente:")
					continue
				}
				switch menu{
				case "1":
					envioRegistro("pymes.csv")
				case "2":
					ShowSeguimiento()
				case "3":
					break
				default:
					fmt.Print("\nFormato u opción no válida, pruebe nuevamente:\n\n")
					continue
				}
			}			
		default:
			fmt.Print("\nFormato u opción no válida, pruebe nuevamente:\n\n")
			continue
		} 


	}
}
