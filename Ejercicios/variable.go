package main 

import "fmt"

func main() {
	
	var numero int //Declaracion de variable
	numero = 17 //Asignandole valor a la variable
	fmt.Println(numero) //Imprimiendo variable en pantalla
	numero = 40 //Cambiando valor
	fmt.Println(numero) //Imprimiendo variable
	//var nombre string 
	nombre := "Angel" //Declaracion de nuevas variables
	nombre = "Miguel" //Asignando nueva variable
	nombre, numero = "Pasillas", 1000 //Asignando valores al mismo tiempo
	nombre2 := "Luis" //Declaracion de nuevas variables
	nombre, nombre2 = nombre2, nombre //Asignando valores al mismo tiempo
	fmt.Println(nombre, numero)
	fmt.Println(nombre, nombre2)

}