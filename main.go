package main

import (
	"fmt"

	"github.com/pablouribe1987/Gopractice/variables"
)

func main() {
	//fmt.Println("Hola Mundo")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
