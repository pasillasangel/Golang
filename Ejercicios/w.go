package main

import (
	"bufio"     //Abrir, leer archivo
	"fmt"       //Impresion en pantalla
	"io/ioutil" //Abrir, leer el archivo
	"os"        //Abrir el archivo
	"strconv"   //convertir int a string
	"strings"   //Convertir los bytes a strings
)

//Tamaño
const size int = 151

func main() {
	var opcion int
	mostrar := `
    --------------------------------------------
          _  _ ____ ____ _  _ _ _  _ ____ 
          |__| |__| [__  |__| | |\ | | __ 
          |  | |  | ___] |  | | | \| |__] 

    --------------------------------------------
    `

	for opcion < 4 {
		fmt.Println(mostrar)
		fmt.Println("\t1.- Escribir")
		fmt.Println("\t2.- Leer")
		fmt.Println("\t3.- Buscar")
		fmt.Println("\t4.- Salir")
		fmt.Print("\tEliga Opcion: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			ingresar()
		case 2:
			leer()
		case 3:
			buscarDato()
		case 4:
			fmt.Println("\nSaliendo del programa.")
		default:
			fmt.Println("\nEliga una opcion correcta.")
		}
	}
}

func ingresar() {
	//Creando variable
	var guardar int
	var hashingMasUno bool = false

	//Ingresando valor a guardar
	fmt.Print("\nIngrese su numero a guardar: ")
	fmt.Scan(&guardar)

	//asignacion hash
	hashing := (guardar % size) + 2

	//Se repetira hasta que encuentre un lugar para aguardarlo
	for i := 0; hashingMasUno != true; i++ {

		//Buscarlo, si ya existe
		b, _ := encontrar(hashing)

		//si lo encuentra que le sume uno y lo vuelva buscar
		if b {
			//probamos con el siguiente hashing (aumentar 1)
			hashing++
		} else {
			//levantar la bandera
			hashingMasUno = true

			//transformacion de int a string
			save := strconv.Itoa(guardar)
			hash := strconv.Itoa(hashing)

			//Mostrar valores
			fmt.Println("#################")
			fmt.Print("\tValor:", save)
			fmt.Println("\tClave:", hash)
			fmt.Println("#################")

			//mandando los valores a escribir en el archivo
			escribir(save, hash)
		}
	}
}

func encontrar(hh int) (bool, int) {
	//transformar de int a string
	h := strconv.Itoa(hh)

	//variables
	bandera := false
	repeticion := 0

	//abrir archivo
	info, _ := ioutil.ReadFile("output.txt")

	//transformar de byte a string
	dat := string(info[:])

	//separar datos y convertirlo en arreglo
	valores := strings.Fields(dat)

	//buscar el valor entre todo el archivo
	for i := 0; i < len(valores); i = i + 2 {

		if h == valores[i] {
			bandera = true
			repeticion++
		}
	}
	//regresar valores bandera y repeticion de la palabra
	return bandera, repeticion
}

func escribir(s string, h string) {
	//Abrir el  archivo
	fileHandle, _ := os.OpenFile("output.txt", os.O_APPEND, 0666)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	//concatenando el hash y ll dato
	cadenaGuardar := h + " " + s

	//Escribiendo en el archivo
	fmt.Fprintln(writer, "\r") //Saltar Linea
	fmt.Fprintln(writer, cadenaGuardar)
	writer.Flush()
}

func leer() {
	//abrir archivo
	info, _ := ioutil.ReadFile("output.txt")

	//transformar de byte a string
	dat := string(info[:])

	//separar datos y convertirlo en arreglo
	valores := strings.Fields(dat)

	//mensaje
	fmt.Println("\t##### Leyendo Valores ######")

	//buscar el valor entre todo el archivo
	for i := 0; i < len(valores)-1; i = i + 2 {

		fmt.Println("\tHash: " + valores[i] + ".")
		i++
		fmt.Println("\tValor: " + valores[i] + ".")
		i--
		fmt.Println()
	}
}

func buscarDato() {
	//variable
	var buscar int

	//ingresar dato a buscar
	fmt.Print("\nIngresar el dato a clave: ")
	fmt.Scan(&buscar)

	//transformar de int a string
	busc := strconv.Itoa(buscar)

	//mandar al metodo de buscar
	buscarMetodo(busc)
}

func buscarMetodo(busc string) {
	//variables
	var encontrado bool = false
	var c int

	//abrir archivo
	info, _ := ioutil.ReadFile("output.txt")

	//transformar de byte a string
	dat := string(info[:])

	//separar datos y convertirlo en arreglo
	valores := strings.Fields(dat)

	//buscar el valor entre todo el archivo
	for i := 0; i < len(valores)-1; i = i + 2 {
		if busc == valores[i] {
			encontrado = true
			c = i
		}
	}

	if encontrado {
		fmt.Println("\n\t##### Valor Encontrado #####")
		fmt.Println("\tHash: " + valores[c] + ".")
		c++
		fmt.Println("\tValor: " + valores[c] + ".")
		fmt.Println("\n\t############################")
	} else {
		fmt.Println("\n\t##### Valor NO Encontrado #####")
	}
}
