package ejerinterfaces

import (
	"fmt"

	"github.com/pablouribe1987/Gopractice/interfaces"
	"github.com/pablouribe1987/Gopractice/modelos"
)

func Prueba() {

	//	Pedro := new(modelos.Hombre)
	men := modelos.Hombre{
		Nombre: "hola",
		Edad:   20,
	}

	fmt.Printf("Nombre: %s, y edad: %d \n", men.Nombre, men.Edad)
}

func HumanosRespirando(hu interfaces.Humano) {
	//var valor bool
	hu.Respirar()

	//valor = hu.
	//fmt.Println("valor=", hu.)
	fmt.Printf("Soy un/a %s, y estoy  respirando \n", hu.Sexo())
}
