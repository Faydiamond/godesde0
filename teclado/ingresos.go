package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println(" Por favor digite el numero 1 ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(" error al convertir el numero1 ")
		}
	}
	fmt.Println(" Por favor digite el numero 2 ")
	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(" error al convertir el numero2 ")
		}
	}
}
