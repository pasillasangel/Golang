package main

import (
	"fmt"       //Instrucciones basicas
	"io/ioutil" //Permite abrir archivos de texto
	"regexp"    //Permite
	"strings"   //Permite la conversion de strings (byte a string)
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var (
		//contador = 0
		total = 0
	)
	//var id [160]string

	fmt.Println("\t\nEl archivo contiene: \n")

	//Abriendo el archivo

	//"/home/pasillas/work/usr/haa.txt"
	//"/home/pasillas/Escritorio/ha.txt"
	info, errr := ioutil.ReadFile("/home/pasillas/work/usr/archivoPalabras.txt")
	check(errr)

	fmt.Println("	---------------------------------------------------")
	//Imprimiendo cadena en consola desde el archivo original
	fmt.Print(string(info))
	fmt.Println("	---------------------------------------------------")

	s := string(info[:]) //Transformar de Byte a String

	fmt.Println() //Salto de linea
	//=<> Separador de simbolos
	r := strings.NewReplacer("{", " { ", "}", " } ", ";", " ; ", "++", " ++ ", "(", " ( ", ")", " ) ", "--", " -- ", "+", " + ", "-", " - ", "*", " * ", "/", " / ", "%", " % ", "[", " [ ", "]", " ]", "==", " == ", "<=", " <= ", ">=", " >= ", "!=", " != ", "&&", " && ", "||", " || ", "!", " ! ")
	z := r.Replace(s)

	valores := strings.Fields(z) //Separa solo las palabras de los espacios

	fmt.Println(`	  ------------------------------------------------------`)
	fmt.Println(`	/   N   |	TOKEN		|	LEXEMA		\`)
	fmt.Println(`	| ----------------------------------------------------- |`)

	for i := range valores {

		//Tokens comunes
		asignacion, _ := regexp.MatchString("^=$", valores[i])                                                               //Asignacion
		operadoresArit, _ := regexp.MatchString("^[\\+|\\-|\\*|\\/|\\%|\\^]$", valores[i])                                   //Operadores aritmeticos
		operadoresRela, _ := regexp.MatchString("((^\\<$)|(^\\>$)|(^\\=\\=$)|(^\\!\\=$)|(^\\>\\=$)|(^\\<\\=$))", valores[i]) //Opéradores Relacionales
		operadoresLogicos, _ := regexp.MatchString("((^\\&\\&$)|(^\\|\\|$)|(^!$))", valores[i])
		numerico, _ := regexp.MatchString("^[0-9]*$", valores[i])                                   //Numerico
		finInstruccion, _ := regexp.MatchString("^(\\;)$", valores[i])                              //Final de instruccion
		agrupadores, _ := regexp.MatchString("^[(\\[)|(\\])|(\\()|(\\))|(\\{)|(\\})]$", valores[i]) //Agrupador
		incremento, _ := regexp.MatchString("^(\\+\\+)$", valores[i])                               //Incremento
		decremento, _ := regexp.MatchString("^(\\-\\-)$", valores[i])                               //Decremettno
		literal, _ := regexp.MatchString("^\"([a-zA-Z]|[0-9]|[\\s])*\"$", valores[i])               //Literal
		tipoDato, _ := regexp.MatchString("((^caracxp$)|(^cadenaxp$)|(^entxp$)|(^bytexp$)|(^cortoxp$)|(^largoxp$)|(^boolxp$)|(^flotanxp$)|(^flotan64xp$)|(^uentxp$)|(^ubytexp$)|(^ucortoxp$)|(^ulargoxp$)|(^decimalxp$)|(^varxp$)|(^enumxp$)|(^structxp$))", valores[i])
		tipoRefere, _ := regexp.MatchString("((^claseq$)|(^interfaceq$)|(^delegadoq$)|(^dinamicoq$)|(^objetoq$))", valores[i])
		tipoBoleano, _ := regexp.MatchString("((^verdad$)|(^falso$))", valores[i])
		instruccionIteracion, _ := regexp.MatchString("((^parala$)|(^paracadala$)|(^mientrasla$)|(^hacerla$)|(^enla$))", valores[i])
		instriccionSeleccion, _ := regexp.MatchString("((^sixv$)|(^sinoxvsixv$)|(^sinoxv$)|(^cambiarxv$)|(^casoxv$))", valores[i])
		instruccionSalto, _ := regexp.MatchString("((^descansarw$)|(^continuarw$)|(^retornarw$)|(^defectow$)|(^rendimientow$))", valores[i])
		atraparErrores, _ := regexp.MatchString("((^intentarxv$)|(^atraparxv$)|(^finalxv$)|(^lanzarxv$))", valores[i])
		controlAcceso, _ := regexp.MatchString("((^publicz$)|(^privadoz$)|(^internoz$)|(^protegidoz$))", valores[i])
		modificadores, _ := regexp.MatchString("(^abstractom$)|(^asincronom$)|(^eventom$)|(^externom$)|(^nuevom$)|(^anularm$)|(^parcialm$)|(^lecturasolom$)|(^selladom$)|(^estaticom$)|(^noguardarm$)|(^virtualm$)|(^volatilm$)|(^constantem$)", valores[i])
		parametrosMetodos, _ := regexp.MatchString("((^paramet$)|(^referen$)|(^salida$))", valores[i])
		checks, _ := regexp.MatchString("((^checar$)|(^nochecar$))", valores[i])
		revision, _ := regexp.MatchString("((^arreglado$)|(^bloquear$))", valores[i])
		nombresEspacio, _ := regexp.MatchString("((^espacionombre$)|(^usando$)|(^externo$))", valores[i])
		claveConversiones, _ := regexp.MatchString("((^explicitoc$)|(^implicitoc$)|(^operadorc$))", valores[i])
		claveAcceso, _ := regexp.MatchString("((^base$)|(^esta$)|(^nulo$))", valores[i])
		claveContextuales, _ := regexp.MatchString("((^obtener$)|(^global$)|(^parcial$)|(^remover$)|(^poner$)|(^evaluar$)|(^donde$)|(^anadir$))", valores[i])
		claveOperadores, _ := regexp.MatchString("((^como$)|(^esperar$)|(^es$)|(^tamañode$)|(^tipode$)|(^stackloco$))", valores[i])
		claveConsulta, _ := regexp.MatchString("((^desde$)|(^donde$)|(^seleccionar$)|(^grupo$)|(^dentro$)|(^ordenarpor$)|(^unirse$)|(^dejar$)|(^en$)|(^sobre$)|(^igual$)|(^por$)|(^acendente$)|(^decendente$))", valores[i])
		impresion, _ := regexp.MatchString("((^Efscanf$)|(^Efscanln$)|(^Impri$)|(^Imprif$)|(^Impriln$)|(^Escan$)|(^Escanf$)|(^Escanln$)|(^ESprint$)|(^ESprintf$)|(^ESprintln$)|(^Eescan$)|(^EEscanf$))", valores[i])
		indicadoresFormato, _ := regexp.MatchString("((^omitir$)|(^izquierda$)|(^derecha$)|(^interno$)|(^deca$)|(^octa$)|(^hexa$)|(^mostrarbase$)|(^mostrarpunto$)|(^letramayus$)|(^mostrarposi$)|(^cientifico$)|(^arreglado$)|(^enterobuf$)|(^justarcampo$)|(^basecampo$)|(^flotantecampo$))", valores[i])
		funcPrincipal, _ := regexp.MatchString("((^vaciov$)|(^iniciov$)|(^funcionv$))", valores[i])
		palabrasFunciones, _ := regexp.MatchString("((^Contar$)|(^Igualdoble$)|(^Campos$)|(^Funcampo$)|(^tenerprefijo$)|(^Tenersufijo$)|(^Repetir$)|(^Remplazar$)|(^Cortar$))", valores[i])
		funcMatematicas, _ := regexp.MatchString("((^abstr$)|(^facos$)|(^fasin$)|(^fatan$)|(^fcos$)|(^fcosh$)|(^fexp$)|(^flog$)|(^fmax$)|(^fmin$)|(^fpow$)|(^fsign$)|(^ftan$)|(^ftanh$)|(^ftruncate$))", valores[i])

		palabrasBaja, _ := regexp.MatchString("((^factivate$)|(^fdouble$)|(^flocators$)|(^frollback$)|(^fadd$)|(^fdrop$)|(^flock$)|(^froundceiling$)|(^fafter$)|(^fdssize$)|(^flockmax$)|(^frounddown$)|(^falias$)|(^fdynamic$)|(^froundfloor$)|(^fall$)|(^feach$)|(^flong$)|(^froundhalfdown$)|(^feditproc$)|(^floop$)|(^fallow$)|(^fmaintained$)|(^fasensitive$)|(^frows$)|(^frownumber$)|(^froutine$)|(^froundup$)|(^fmaterialized$)|(^fencryption$)|(^fassociate$)|(^fmaxvalued$)|(^fdefinition$)|(^flanguage$)|(^frelease$)|(^fvolumes$)|(^frename$)|(^freset$)|(^fleve$)|(^fctype$))", valores[i])

		identificador, _ := regexp.MatchString("^[a-zA-Z]([a-zA-Z]|[0-9])*", valores[i]) //Identificadores

		if tipoDato {
			fmt.Println(`	|`, total, `	|	TIPO DATO	|	 `+valores[i]+`	|`)
		} else if tipoBoleano {
			fmt.Println(`	|`, total, `	|	TIPO BOOL	|	 `+valores[i]+`	|`)
		} else if instruccionSalto {
			fmt.Println(`	|`, total, `	|	INSTR SALTO	|	 `+valores[i]+`	|`)
		} else if atraparErrores {
			fmt.Println(`	|`, total, `	|	ATRAP ERROR	|	 `+valores[i]+`	|`)
		} else if controlAcceso {
			fmt.Println(`	|`, total, `	|	CONT ACCESO	|	 `+valores[i]+`	|`)
		} else if asignacion {
			fmt.Println(`	|`, total, `	|	ASIGNACION	|	 =	|`)
		} else if operadoresArit {
			fmt.Println(`	|`, total, `	|	OPERA ARIT	|	 `+valores[i]+`		|`)
		} else if operadoresRela {
			fmt.Println(`	|`, total, `	|	OPERA RELA	|	 `+valores[i]+`	|`)
		} else if numerico {
			fmt.Println(`	|`, total, `	|	NUMERICO	|	 `+valores[i]+`	|`)
		} else if finInstruccion {
			fmt.Println(`	|`, total, `	|	FIN INSTRUC	|	 ;	|`)
		} else if agrupadores {
			fmt.Println(`	|`, total, `	|	AGRUPADORES	|	 `+valores[i]+`	|`)
		} else if incremento {
			fmt.Println(`	|`, total, `	|	INCREMENTO	| 	 ++	|`)
		} else if decremento {
			fmt.Println(`	|`, total, `	|	DECREMENTO	| 	 --	|`)
		} else if literal {
			fmt.Println(`	|`, total, `	|	LITERAL		|	 `+valores[i]+`	|`)
		} else if modificadores {
			fmt.Println(`	|`, total, `	|	MODIFICA	|	 `+valores[i]+`	|`)
		} else if tipoRefere {
			fmt.Println(`	|`, total, `	|	T REFERENCIA	|	 `+valores[i]+`	|`)
		} else if instriccionSeleccion {
			fmt.Println(`	|`, total, `	|	INSTR SELEC	|	 `+valores[i]+`	|`)
		} else if instruccionIteracion {
			fmt.Println(`	|`, total, `	|	INSTR ITERA	|	 `+valores[i]+`	|`)
		} else if parametrosMetodos {
			fmt.Println(`	|`, total, `	|	PARA METOD	|	 `+valores[i]+`	|`)
		} else if checks {
			fmt.Println(`	|`, total, `	|	CHECKS		|	 `+valores[i]+`	|`)
		} else if revision {
			fmt.Println(`	|`, total, `	|	REVISION	|	 `+valores[i]+`	|`)
		} else if nombresEspacio {
			fmt.Println(`	|`, total, `	|	NOMB ESPAC	|	 `+valores[i]+`	|`)
		} else if claveOperadores {
			fmt.Println(`	|`, total, `	|	CLAVE OPERA	|	 `+valores[i]+`	|`)
		} else if claveConversiones {
			fmt.Println(`	|`, total, `	|	CLAVE CONVER	|	 `+valores[i]+`	|`)
		} else if claveAcceso {
			fmt.Println(`	|`, total, `	|	CLAVE ACCESO	|	 `+valores[i]+`	|`)
		} else if claveContextuales {
			fmt.Println(`	|`, total, `	|	CLAVE CONTEXTUA	|	 `+valores[i]+`	|`)
		} else if claveConsulta {
			fmt.Println(`	|`, total, `	|	CLAVE CONSULTA	|	 `+valores[i]+`	|`)
		} else if indicadoresFormato {
			fmt.Println(`	|`, total, `	|	INDICA FORMATO	|	 `+valores[i]+`	|`)
		} else if impresion {
			fmt.Println(`	|`, total, `	|	IMPRESION	|	 `+valores[i]+`	|`)
		} else if palabrasFunciones {
			fmt.Println(`	|`, total, `	|	PALABRAS FUNC	|	 `+valores[i]+`	|`)
		} else if operadoresLogicos {
			fmt.Println(`	|`, total, `	|	OPERA LOGICOS	|	 `+valores[i]+`		|`)
		} else if funcPrincipal {
			fmt.Println(`	|`, total, `	|	FUNC PRINCIPAL	|	 `+valores[i]+`	|`)
		} else if funcMatematicas {
			fmt.Println(`	|`, total, `	|	FUNC MATEMAT	|	 `+valores[i]+`	|`)
		} else if palabrasBaja {
			fmt.Println(`	|`, total, `	|	PALABRAS FUNC	|	 `+valores[i]+`	|`)
		} else if identificador {
			fmt.Println(`	|`, total, `	|	ID		|	`+valores[i]+`	|`)
			//id[contador] = valores[i]
			//contador++
		} else {
			fmt.Println(`	|`, total, `	|	ERROR		|	 `+valores[i]+`	|`)
		}

		total++ //Aumentador el contador
	}

	fmt.Println(`	\ ------------------------------------------------------ /`)
	fmt.Println()
	/*
			fmt.Println("\n\t\tTABLA DE SIMBOLOS")
			fmt.Println(`	  -----------------------------`)
			fmt.Println(`	/	ID	|	 VALOR	\`)
			fmt.Println(`	| ----------------------------- |`)
			for e := 0; e < contador; e++ {
				fmt.Println(`	|	`, e, `	|	`+id[e]+`	|`)
			}

		fmt.Println(`	\ ----------------------------- /`)
		fmt.Println()
	*/
}
