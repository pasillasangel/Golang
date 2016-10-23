package main 

import (
	"fmt"
	"strconv"
)

func main() {
	var cadena string
	cadena = "HOLA MUNDILLO"

	fmt.Println(cadena) //Imprimir cadena
	fmt.Println(len(cadena)-1) //IMprime el numero de caracteres de la cadena

	fmt.Println(cadena[2]) //Imprime el valor de la cadena pero en UTF-8

	fmt.Println(cadena[0:4]) //Imprimir parte de la variable

	cadena = cadena + " OTRA VEZ"
	fmt.Println(cadena)

	cadena += " siiii"
	fmt.Println(cadena)

	//STRING LITERAL
	cadena = `
		<!DOCTYPE html>
		<html>
			<head>
				<title></title>
			</head>
			<body>

			</body>
		</html>
	`
	fmt.Println(cadena)

	cadena = "Hola \"Mundo\""
	fmt.Println(cadena)

	//CONCATENACION ******************************************
	suma := 5
	cadena = "La suma es " + strconv.Itoa(suma)
	fmt.Println(cadena)
}