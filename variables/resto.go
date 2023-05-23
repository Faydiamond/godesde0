package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Salario float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Fabian"
	Estado = true
	Salario = 3000000
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Salario)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
