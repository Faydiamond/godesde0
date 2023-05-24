package main

import (
	"fmt"
	"runtime"

	"github.com/Faydiamond/godesde0/ejercicios"
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

	if os := runtime.GOOS; os == "windows" {
		fmt.Println("Su sistema operativo es  Windows")
	}

	switch oss := runtime.GOOS; oss {
	//case "windows":
	//	fmt.Println("Es un Windows")
	case "ubuntu":
		fmt.Println("Es un ubuntu")
	default:
		fmt.Printf("%s  ", oss)
	}
	fmt.Println("////////////////")

	valor, res := ejercicios.Conversiones("150")
	fmt.Println("El valor es: ", valor)
	fmt.Printf("El resultado  es: %s ", res)
}
