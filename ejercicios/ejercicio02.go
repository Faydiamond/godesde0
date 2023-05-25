package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//var err error

func Multiplicar() string {
	fmt.Println(" Por favor digite el numero de la tabla del multiplicar que desea conocer ")
	var texto string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {

		numero1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Por favor ingresa un nuemero ")
			Multiplicar()
		} else {
			for i := 1; i <= 10; i++ {
				texto += fmt.Sprintf("%d  X %d = %d \n", numero1, i, numero1*i)
			}
		}
	}
	//fmt.Println(texto)
	return texto
}
