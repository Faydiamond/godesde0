package ejercicios

import (
	"strconv"
)

func Conversiones(cadena string) (int, string) {

	if res, _ := strconv.Atoi(cadena); res > 100 {
		return res, "Es mayor a 100"
	} else {
		return res, "Es menor a 100"
	}
}
