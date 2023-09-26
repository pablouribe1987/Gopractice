package main

import (
	"fmt"

	e "github.com/pablouribe1987/Gopractice/ejerinterfaces"
	"github.com/pablouribe1987/Gopractice/modelos"
)

func main() {
	//fmt.Println("Hola Mundo")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows")

	} else {
		fmt.Println("Esto es Windows")

	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "Darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/
	//funciones.Calculos()
	Pedro := new(modelos.Hombre)

	e.HumanosRespirando(Pedro)
	fmt.Println("valor=", Pedro.Respirando)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)

	e.Prueba()
}
