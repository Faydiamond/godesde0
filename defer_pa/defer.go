package defer_

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Primer mensaje")
	defer fmt.Println("Este mensaje final")

	fmt.Println("Segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			//log siempre antepone la fecha y la hora de la ejecucion
			log.Fatalf("Ocurrio un error que genero panic \n %v ", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
