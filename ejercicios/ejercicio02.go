package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//var err error

func Multiplicar() {
	fmt.Println(" Por favor digite el numero de la tabla del multiplicar que desea conocer ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Por favor ingresa un nuemero ")
			Multiplicar()
		} else {
			for i := 1; i <= 10; i++ {
				fmt.Printf("%d ", numero1*i)
			}
		}
	}
}
