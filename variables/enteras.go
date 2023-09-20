package variables
import "fmt"

//para que la funcion sea publica la primera letra tiene que ser en mayuscula
func MuestroEnteros(){
	var intcomun int
	intde32:= int32(10)
	intde64:= int64(100)

	fmt.Println("intcomun =",intcomun)
	fmt.Println("intde32 =",intde32)
	fmt.Println("intde64 =",intde64)
}