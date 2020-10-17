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
	dire = "localhost:50051"
)

type Envio struct{
	id string
	producto string
	valor int64
	tienda string
	destino string
	prioritario int64 //{normal, prioritario, retail}={0,1,2}
}
var tiempopaquetes string
var tiempopaquetesint int
func envioRegistro(archivo string)(){
	conn, err := grpc.Dial(dire, grpc.WithInsecure())
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
	var lista_envios []Envio

	for {
		record, err := reader.Read()

		if err == io.EOF { // cuando termina de leer el archivo
			break 
		}
		
		if record[0] == "id"{
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
		lista_envios = append(lista_envios, envio)
	}

	//fmt.Println(lista_envios)
	//Envio de retail
	
	for i := 0; i < len(lista_envios); i++ {
		
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
		
		r, err := c.Envio(ctx, envio)
		time.Sleep(time.Duration(tiempopaquetesint*1000)*time.Millisecond)
		if err != nil {
			log.Fatalf("Requ err: %v", err)
		}
		
		
		fmt.Println("Codigo de seguimiento: "+r.GetMsg())

	}	
}

func ShowSeguimiento (){
	conn, err := grpc.Dial(dire, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn err: %v", err)

	}
	defer conn.Close()
	fmt.Println(" Ingrese codigo de seguimiento:  ")
	var seguimiento string
	fmt.Scanln(&seguimiento)
	c := cl.NewEnvioServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	seg := &cl.SeguimientoRequest{
		Seguimiento: seguimiento,
	}

	r, _ := c.Seguimiento(ctx, seg)

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
		fmt.Print("Menu Cliente \n1.-Ingreso de retail\n2.-Ingreso de Pyme\n3.-Realizar Seguimiento\nIngrese opción:")
		_,err:=fmt.Scanln(&menu)
		if err!=nil{
			fmt.Print("\nFormato de ingreso no válido, pruebe nuevamente:")
			continue
		}
		switch menu{
		case "1":
			envioRegistro("retail.csv")
		case "2":
			envioRegistro("pymes.csv")
		case "3":
			ShowSeguimiento()
		default:
			fmt.Print("\nFormato u opción no válida, pruebe nuevamente:\n\n")
			continue
		} 


	}
	//fmt.Println("Usuario del sistema PrestigioExpress, por favor ingresar el tiempo de envio de pedidos a logistica (en segundos): ")
	//var tiempo int64
	//fmt.Scanln(&tiempo)  
	// Creando conexion TCP
	

	// Lectura de Retail
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

}
