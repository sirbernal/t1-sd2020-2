package main

import (
	"fmt"
)

func main() {
	var menu string
	
	LoopMain:
		for {
		fmt.Print("\nMenu Principal \n1.-Modo retail\n2.-Modo Pyme\n3Terminar Aplicacion Cliente\nIngrese opción:")
		_,err:=fmt.Scanln(&menu)
		if err!=nil{
			fmt.Print("\nFormato de ingreso no válido, pruebe nuevamente:")
			continue
		}
		switch menu{
		case "1":
			var submenu1 string
			submenuretail:
				for {
					fmt.Print("\nMenu Retail \n1.-Enviar Registro\n2.-Volver\nIngrese opción:")
					_,err:=fmt.Scanln(&submenu1)
					if err!=nil{
						fmt.Print("\nFormato de ingreso no válido, pruebe nuevamente:")
						continue
					}
					switch submenu1{
					case "1":
						fmt.Println("Opcion 1.1")
					case "2":
						fmt.Println("Opcion 1.2")
						break submenuretail
					default:
						fmt.Println("Opcion no válida")
					}
				
				}
		case "2":
			var submenu2 string
			submenupyme:
				for {
					fmt.Print("\n\n\nMenu Pyme \n1.-Enviar Registro\n2.-Realizar Seguimiento\n3.-Volver\nIngrese opción:")
					_,err:=fmt.Scanln(&submenu2)
					if err!=nil{
						fmt.Print("\nFormato de ingreso no válido, pruebe nuevamente:")
						continue
					}
					switch submenu2{
					case "1":
						fmt.Println("Enviando registro")
					case "2":
						fmt.Println("Haciendo seguimiento")
					case "3":
						fmt.Println("Regresando Atras")
						break submenupyme
					default:
						fmt.Println("Opcion no válida")
					}
				
				}
			fmt.Println("Envio de registros completados!")
		case "3":
			fmt.Println("Cerrando App")
			break LoopMain
		default:
			fmt.Print("\nFormato u opción no válida, pruebe nuevamente:\n\n")
			continue
		} 


	}
}
