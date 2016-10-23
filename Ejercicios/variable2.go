package main 

import "fmt"

func main() {

	//Variable sin especificar su TIPO
	var numero = 20

	//Declaracion de variables
	var (
		a = 5
		b = 10
		c = 15
	)
	//Concatenacion de valores
	fmt.Println("Los numeros son:", a , b , c, numero, ". Sumados dan: ", (a+b+c+numero))

	suma := 5

	fmt.Println("Los suma de 3 mas 2 es: " + suma)
}