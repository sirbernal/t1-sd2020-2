package main
import (
	"context"
	"log"
	"fmt"
	"time"
	"reflect"
	"strings"
	pb "github.com/sirbernal/t1-sd2020-2/proto/camion_logistica"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"os"
	"encoding/csv"
)

const (
	dire = "10.10.28.82:50051"   //Direccion maquina logistica para hacer coneccion grpc
)

type Envio struct{
	idPaquete string
	seguimiento int64
	tipo int64 //{normal, prioritario, retail}={0,1,2}
	valor int64
	intentos int64
	estado int64 //{En bodega, En camino, Recibido o No Recibido}={0,1,2,3}
	fecha_entrega time.Time
}
//listas que guardan los datos de cada camion para generar el archivo resumen
var camionretail1 [][] string
var camionretail2 [][] string
var camionnormal [][] string
//variables (con auxiliares int) para guardar tiempos de espera
var tiempocamiones string
var tiempocamionesint int
var tiempoentrega string
var tiempoentregaint int
func updateCamion(result [6]Envio){    //Funcion que guarda los registros de los envios de los paquetes en archivos csv
	for i,pack := range result{
		if reflect.DeepEqual(pack,Envio{estado:1}){
			continue
		}
		var tipo string
		var archivo string
		var fecha string
		narch:=i/2
		switch pack.tipo{
		case 0:
			tipo="normal"
		case 1:
			tipo="prioritario"
		case 2:
			tipo="retail"
		default:
			continue
		}
		if pack.fecha_entrega==(time.Time{}){
			fecha="0"
		}else{
			fecha=pack.fecha_entrega.Format("02-01-2006 15:04:05")
		}

		conn, err := grpc.Dial(dire, grpc.WithInsecure())
		
		if err != nil {
			log.Fatalf("Conn err: %v", err)
		}

		c := pb.NewCamionServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		seg := &pb.DatosRequest{
			Name: pack.idPaquete+","+strconv.FormatInt(pack.valor,10)+","+strconv.FormatInt(pack.seguimiento,10),
		}

		r, _ := c.DatosCamion(ctx, seg)

		resp := strings.Split(r.GetDato(),",")

		linea:= []string{pack.idPaquete,tipo,
			strconv.FormatInt(pack.valor,10), resp[0], resp[1],
			strconv.FormatInt(pack.intentos,10),fecha}
		switch narch{
		case 0:
			archivo="camionretail1.csv"
			camionretail1=append(camionretail1,linea)
			file,err:= os.OpenFile(archivo,os.O_CREATE|os.O_WRONLY,0777)
			defer file.Close()
			if err !=nil{
				os.Exit(1)
			}
			csvWriter:= csv.NewWriter(file)
			csvWriter.WriteAll(camionretail1)
			csvWriter.Flush()
		case 1:
			archivo="camionretail2.csv"
			camionretail2=append(camionretail2,linea)
			file,err:= os.OpenFile(archivo,os.O_CREATE|os.O_WRONLY,0777)
			defer file.Close()
			if err !=nil{
				os.Exit(1)
			}
			csvWriter:= csv.NewWriter(file)
			csvWriter.WriteAll(camionretail2)
			csvWriter.Flush()
		case 2:
			archivo="camionnormal.csv"
			camionnormal=append(camionnormal,linea)
			file,err:= os.OpenFile(archivo,os.O_CREATE|os.O_WRONLY,0777)
			defer file.Close()
			if err !=nil{
				os.Exit(1)
			}
			csvWriter:= csv.NewWriter(file)
			csvWriter.WriteAll(camionnormal)
			csvWriter.Flush()
		default:
			continue
		} 

	}
	

}
func menorEnvio(x Envio, y Envio)(Envio, Envio){ //ordena entre dos envios de mayor a menor segun costo
	if x.valor>y.valor{
		return x,y
	}else{
		return y,x
	}
}
func viaje(env [2]Envio)[2]Envio{  //funcion invocada al hacer una simulacion, lo que realizara en teoria el viaje
	void:= Envio{             // por camion (osea 2 paquetes) y retornara los resutados de los paquetes tras el viaje
		estado:1,         // Es aqui donde funcionan las probabilidades mencionadas en el enunciado y el hacer intentos
	}                         // de entrega
	for i,pac := range env{
		if reflect.DeepEqual(pac,void){ //si el paquete es vacio (no existe) evita realizar el envio
			continue
		}
		fmt.Println("Enviando paquete"+pac.idPaquete+" ,Intento: "+strconv.FormatInt(pac.intentos+1,10))
		time.Sleep(time.Duration(tiempoentregaint*1000)*time.Millisecond)
		for x:=0;x<4;x++{
			if x==3{
				fmt.Println("Límite de envios superado\n")
				pac.estado=3
				break
			}
			if pac.tipo<2 && int64(x)*10>pac.valor{
				fmt.Println("Límite de costo superado\n")
				pac.estado=3
				break
			}
			pac.intentos++
			rand.Seed(time.Now().UnixNano()) 
			probabilidad:= rand.Intn(101)
			if probabilidad<80{
				fmt.Println("Pedido Entregado!\n")
				pac.estado=2
				pac.fecha_entrega=time.Now()
				break
			}else{
				fmt.Println("Entrega fallida... procesando nuevo intento de entrega...")
			}
		}
		env[i]=pac
	}
	return env
}
func simularEnvio(env [6]Envio)[6]Envio{    // Estan función recibe el arreglo de 6 paquetes, para simular los envios en cada camion y retornará el arreglo con los resultados de la simulacion
	fmt.Println("Inicializando envio de camiones...")
	fmt.Println(env)
	e0,e1 := menorEnvio(env[0],env[1])  //Ordena en cada camion los paquetes en funcion del valor (mayor primero)
	e2,e3 := menorEnvio(env[2],env[3])
	e4,e5 := menorEnvio(env[4],env[5])
	cam1 :=[2]Envio{e0,e1} //camiones ya ordenado
	cam2 :=[2]Envio{e2,e3}
	cam3 :=[2]Envio{e4,e5} 
	fmt.Println("Camión Retail 1")
	cam1 = viaje(cam1) //simula los viajes de cada camion por separado
	fmt.Println("Camión Retail 2")
	cam2 = viaje(cam2)
	fmt.Println("Camión Normal")
	cam3 = viaje(cam3)
	resultado := [6]Envio{cam1[0], cam1[1], cam2[0], cam2[1], cam3[0], cam3[1]} //guarda el resultado de todos los camiones
	updateCamion(resultado) //guarda los registros de cada camion en memoria para actualizar los archivos de cada uno
	return resultado //retorna el resultado para enviar a logistica

}

func main() {
	// En estas primeras lineas se pide los tiempos de espera y demora de los camiones
	fmt.Print("Inicializando servicio de reparto...Listo!\nIngrese tiempo (segundos) de espera de camiones:")
	for {
		fmt.Scanln(&tiempocamiones)
		intmode, err:=strconv.Atoi(tiempocamiones)
		if err!=nil{
			fmt.Print("Formato no válido, ingrese nuevamente:")
			continue
		}
		tiempocamionesint=intmode
		break
	}
	fmt.Print("Ingrese tiempo (segundos) de demora de entregas:")
	for{
		fmt.Scanln(&tiempoentrega)
		intmode, err:=strconv.Atoi(tiempoentrega)
		if err!=nil{
			fmt.Print("Formato no válido, ingrese nuevamente:")
			continue
		}
		tiempoentregaint=intmode
		break
	}

	// Creando conexion TCP
	conn, err := grpc.Dial(dire, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Conn err: %v", err)
	}
	var envios [6]Envio
	client := pb.NewCamionServiceClient(conn)  // Se define la conexion con el servicio definido en el proto 
	stream, err := client.Camion(context.Background()) // Como se uso un stream, se inicializa aca
	waitc := make(chan struct{})  // Se define una variable de canal que nos permite mantener las goroutines activas durante 
	// todo el proceso

	// La logica del camion primero es enviar un paquete "vacio" para avisarle a logistica que este servicio esta activo
	// para hacer andar los camiones

	msg := &pb.CamionRequest{IdPaquete : "",     
		Seguimiento : -1,
		Tipo :0,
		Valor :  0,
		Intentos : 0,
		Estado : 0,}
	go func() {    // Enviamos nuestro primer mensaje que mencionamos anteriormente en esta goroutine
		for i := 0; i < 1; i++ {  
		stream.Send(msg) }
		
	}()
	go func() {  // Esta goroutine estara encargada basicamente de recibir paquetes de logistica y hacer todo
		npack:= 0
		var resultado [6]Envio // Para guardar los paquetes que se reciben
		for {     // Recibimos los paquetes de logistica
			resp, err := stream.Recv()
			paquete := Envio{          // los paquetes los guardamos en esta clase
				idPaquete : resp.IdPaquete,
				seguimiento : resp.Seguimiento,
				tipo :resp.Tipo,
				valor :  resp.Valor,
				intentos : resp.Intentos,
				estado : 1,
			}
			envios[npack]=paquete  // Se guarda cada paquete en el arreglo de arriba
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			npack++
			if npack==6{  // Al recibir todos los 6 paquetes correspondientes a 2 de cada camion, se simula los envios
				if !(reflect.DeepEqual(envios,[6]Envio{Envio{estado:1},Envio{estado:1},Envio{estado:1},Envio{estado:1},Envio{estado:1},Envio{estado:1}})){
					resultado = simularEnvio(envios)  // Simulamos envios, definido arriba. Recibiremos un arreglo con los resultados
					for _, pack := range resultado{ // Paso siguiente: enviar devuelta cada resultado de los envios de los paquetes
						msg2 := &pb.CamionRequest{    // definimos lo que vamos a enviar por cada paquete
							IdPaquete: pack.idPaquete,   
							Seguimiento: pack.seguimiento,
							Tipo: pack.tipo,        
							Valor: pack.valor,       
							Intentos: pack.intentos,    
							Estado: pack.estado }
						stream.Send(msg2) // enviamos
						// Tiempo de espera definido arriba x el usuario
					}
					npack++
				}
				
			}
			if npack==7{    // Cuando terminamos de hacer el recebimieno, simulacion y entrega de resultados de los 6 
				stream.Send(msg)  //paquetes, nuevamente usamos nuestro mensaje inicial para avisar a logistica que podemos
				npack=0        // recibir 6 paquetes mas
			}
			time.Sleep(time.Duration(tiempocamionesint*1000)*time.Millisecond) // Tiempo de espera que se definio al principio
			//fmt.Println(envios)
		}


	}()
	<-waitc  // esto en teoria permite que las goroutines funcionen todo el tiempo (segun internet)
	stream.CloseSend()
}


