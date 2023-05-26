package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MinombreLento(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "") //por cada letra
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

	canal1 <- true //le asigna un valor explisitamente

}
