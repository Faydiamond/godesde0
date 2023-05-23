package main

import (
	"fmt"

	"github.com/Faydiamond/godesde0/variables"
)

func main() {
	variables.MostrarEnteros()
	fmt.Println("////////////////")
	variables.RestoVariables()
	fmt.Println("////////////////")
	estado, resultado := variables.ConviertoaTexto(2500)
	fmt.Println("Estado: ", estado)
	fmt.Println("Resultado: ", resultado)
}
