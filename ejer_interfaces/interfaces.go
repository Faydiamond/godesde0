package ejerinterfaces

import (
	"fmt"

	"github.com/Faydiamond/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf(" Soy un %s   y estoy respirando ", hu.Genero())
}
