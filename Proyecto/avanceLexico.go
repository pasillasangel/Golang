package main

import (
	"bufio"     //Libreias para detener la pantalla
	"fmt"       //Comando basicos
	"io/ioutil" //Permite abrir archivos de texto
	"os"        //Libreias para detener la pantalla
	"regexp"    //Expresiones regulares basados en PERL
	"strings"   //Permite la conversion de strings (byte a string)
)

func main() {

	//Variables
	var continuar string //Variable para continuar en el programa
	var cadena string    //Guarda la cadena ingresada por el usuario
	var opcion uint8     //Guarga la opcion elegida por el usuario
	var (                //Enteros para contadores de variables
		i1     = 0
		i2     = 0
		i3     = 0
		i4     = 0
		i5     = 0
		itotal = 0
	)
	var ( //Arreglos para almacenar las cadenas provenientes del archivo
		p1 [20]string
		p2 [20]string
		p3 [20]string
		p4 [20]string
		p5 [20]string
	)

	//Diseño
	asteriscos := `	****************************************`
	pantallaInicio0 := `





		****************************************
		*                                      *
		*   Identificar el grupo de la cadena: *
		*     1) Desde Una Cadena              *
		*     2) Desde un Archivo              *
		*                                      *
		****************************************`
	pantallaInicio1 := `





		****************************************
		*               OPCION 1               *
		*     Reconocer una cadena entre un:   *
		*     -Identificador                   *
		*     -Numerico                        *
		*     -Correo Electronico              *
		*     -URL                             *
		*                                      *
		****************************************`
	pantallaInicio2 := `





		****************************************
		*               OPCION 2               *
		*     Reconocer una cadena desde un    *
		*     un ARCHIVO entre un:             *
		*     -Identificador                   *
		*     -Numerico                        *
		*     -Correo Electronico              *
		*     -URL                             *
		*                                      *
		****************************************`

	continuar = "s" //Asignando valor a la variable para iniciar en el ciclo

	for (continuar == "s") || (continuar == "S") { //Si es "s" o "S" seguir en el ciclo

		//Variables en 0
		i1 = 0
		i2 = 0
		i3 = 0
		i4 = 0
		i5 = 0
		itotal = 0
		opcion = 0
		cadena = ""

		//Regresar a cero
		for z := 0; z < 20; z++ {
			p1[z] = ""
			p2[z] = ""
			p3[z] = ""
			p4[z] = ""
			p5[z] = ""
		}

		fmt.Println(pantallaInicio0)
		fmt.Println(asteriscos)
		fmt.Print("\tELige una opcion: ")
		fmt.Scan(&opcion)
		fmt.Println(asteriscos + "\n")

		switch opcion {
		case 1: //Opcion: desde cadena
			fmt.Println(pantallaInicio1)
			fmt.Println(asteriscos)
			fmt.Print("\tIngrese una cadena: ")
			fmt.Scan(&cadena) //Ingresando la candena identificar
			fmt.Println(asteriscos + "\n")

			//PATRONES
			//Email: Letra o digito seguido de letra o digito o punto o guion o guion bajo seuido de una sola arroba seguido de letra o numero seguida de un solo punto seguido de letra
			patron1, _ := regexp.MatchString("((^[_a-zA-Z-0-9-]+([_.-][a-zA-Z0-9]+)*)@[a-zA-Z0-9-]+(.[a-zA-Z0-9-]+)*(.[a-z])$)", cadena)
			//Numero: Digito seguido de un digito
			patron2, _ := regexp.MatchString("(^[0-9]*$)", cadena)
			//Identificador: Letra seguida de letra o numero incluyendo la cadena vacia
			patron3, _ := regexp.MatchString("(^[a-zA-Z]([a-zA-Z]|[0-9])*$)", cadena)
			//URL: Iniciar con http o https seguido de letra(s) o numero(s) seguido de punto seguido de letra(s)....
			patron4, _ := regexp.MatchString("(^(https?://)?([0-9a-zA-Z.-]+).([a-zA-Z]{2,6})([//a-zA-Z=.-]*)*/?$)", cadena)

			//Mostrando Resultado
			fmt.Print(asteriscos)
			if patron1 {
				fmt.Print("\n\tLa Cadena es un CORREO ELECTRONICO.\n")
			} else if patron2 {
				fmt.Print("\n\tLa Cadena es NUMERIC0.\n")
			} else if patron3 {
				fmt.Print("\n\tLa Cadena es un IDENTIFICADOR.\n")
			} else if patron4 {
				fmt.Print("\n\tLa Cadena es un URL.\n")
			} else {
				fmt.Print("\n\tLa Cadena NO COINCIDE CON NINGUN GRUPO.\n")
			}
			fmt.Println(asteriscos + "\n")

		case 2: //Opcion: desde Archivo
			fmt.Println(pantallaInicio2)
			fmt.Println(asteriscos)
			fmt.Print("\tPresione < ENTER > para leer el archivo de texto: ")
			bufio.NewReader(os.Stdin).ReadBytes('\n') //Detener pantalla
			fmt.Println(asteriscos + "\n\n")

			//*****Leer archivo dsde su direccion
			// Perhaps the most basic file reading task is
			// slurping a file's entire contents into memory.
			// ("/home/pasillas/Escritorio/ha.txt")
			info, errr := ioutil.ReadFile("/home/pasillas/Escritorio/archivoPalabras.txt")
			check(errr)

			fmt.Println(asteriscos)
			fmt.Println(asteriscos)
			fmt.Print(string(info)) //Imprimiendo cadena en consola desde el archivo original
			fmt.Println(asteriscos)
			fmt.Println(asteriscos + "\n")

			s := string(info[:]) //Transformar de Byte a String**

			//Cortar valores string donde se encuentre un espacio " "
			//nuevosStrings := strings.Split(s, " ")

			//Separa solo las palabras de todo lo demas
			//nuevosStrings := strings.FieldsFunc(s, f)

			//Separa solo las palabras de los espacios
			nuevosStrings := strings.Fields(s)

			fmt.Println(asteriscos)
			fmt.Print("\tPresione <Enter> para separar las palabras y encontrarles grupo: ")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			fmt.Println(asteriscos + "\n")

			// Evaluando todos los valores
			for i := range nuevosStrings {

				//PATRONES
				//Correo: Letra o digito seguido de letra o digito o punto o guion o guion bajo seuido de una sola arroba seguido de letra o numero seguida de un solo punto seguido de letra
				patron11, _ := regexp.MatchString("((^[_a-zA-Z-0-9-]+([_.-][a-zA-Z0-9]+)*)@[a-zA-Z0-9-]+(.[a-zA-Z0-9-]+)*(.[a-z])", nuevosStrings[i])
				//Numero: Digito seguido de un digito
				patron22, _ := regexp.MatchString("(^[0-9]*$)", nuevosStrings[i])
				//Identificador: Letra seguida de letra o numero incluyendo la cadena vacia
				patron33, _ := regexp.MatchString("(^[a-zA-Z]([a-zA-Z]|[0-9])*$)", nuevosStrings[i])
				//URL: Iniciar con http o https seguido de letra(s) o numero(s) seguido de punto seguido de letra(s)....
				patron44, _ := regexp.MatchString("(^(https?://)?([0-9a-zA-Z.-]+).([a-zA-Z]{2,6})([//a-zA-Z=.-]*)*/?$)", nuevosStrings[i])

				//patron55, _ := regexp.MatchString("^(=)$", nuevosStrings[i])

				//Mostrando Resultado
				if patron11 { //Correo
					p1[i1] = nuevosStrings[i] //Guardar valor en el arreglo
					i1++                      //Aumentar variable
				} else if patron22 { //Numerico
					p2[i2] = nuevosStrings[i] //Guardar valor en el arreglo
					i2++                      //Aumentar variable
				} else if patron33 { //Identificador
					p3[i3] = nuevosStrings[i] //Guardar valor en el arreglo
					i3++                      //Aumentar variable
				} else if patron44 { //URL
					p4[i4] = nuevosStrings[i] //Guardar valor en el arreglo
					i4++                      //Aumentar variable
				} else { //Ningun grupo
					p5[i5] = nuevosStrings[i] //Guardar valor en el arreglo
					i5++                      //Aumentar variable
				}
				//Contar todas las palabras
				itotal++
			}

			fmt.Println(asteriscos)
			fmt.Println("\tTotal de Cadenas encontradas: ", itotal)
			fmt.Println(asteriscos + "\n\n")

			fmt.Println(asteriscos)
			fmt.Println("\tCadenas de Correos encontrados: ", (i1))
			fmt.Println(asteriscos)
			for a := 0; a < i1; a++ {
				fmt.Println("\t" + p1[a])
			}
			fmt.Println(asteriscos + "\n\n")

			fmt.Println(asteriscos)
			fmt.Println("\tCadenas de Numeros encontrados: ", (i2))
			fmt.Println(asteriscos)
			for b := 0; b < i2; b++ {
				fmt.Println("\t" + p2[b])
			}
			fmt.Println(asteriscos + "\n\n")

			fmt.Println(asteriscos)
			fmt.Println("\tCadenas de Identificadores encontradas: ", (i3))
			fmt.Println(asteriscos)
			for c := 0; c < i3; c++ {
				fmt.Println("\t" + p3[c])
			}
			fmt.Println(asteriscos + "\n\n")

			fmt.Println(asteriscos)
			fmt.Println("\tCadenas de URL's encontradas: ", (i4))
			fmt.Println(asteriscos)
			for d := 0; d < i4; d++ {
				fmt.Println("\t" + p4[d])
			}
			fmt.Println(asteriscos + "\n\n")

			fmt.Println(asteriscos)
			fmt.Println("\tCadenas Desconocidas encontrados: ", (i5))
			fmt.Println(asteriscos)
			for e := 0; e < i5; e++ {
				fmt.Println("\t" + p5[e])
			}
			fmt.Println(asteriscos + "\n\n\n")

		default: //Opcion: por defector (error)
			fmt.Print(asteriscos)
			fmt.Print("\n\t*******Eliga una opcion correcta.*******\n")
			fmt.Println(asteriscos + "\n")

		}

		//Preguntar si quiere volver a intentar
		fmt.Println(asteriscos)
		fmt.Print("\t¿Desea realizar otro intento? (s/n): ")
		fmt.Scan(&continuar)
		fmt.Println(asteriscos + "\n")
	}
}

//*****Metodo para error en el archivo txt (si llegara a haber)
// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
