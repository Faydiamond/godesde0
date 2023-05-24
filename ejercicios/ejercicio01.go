package ejercicios

import (
	"strconv"
)

func Conversiones(cadena string) (int, string) {

	res, error := strconv.Atoi(cadena)
	if error != nil {
		return 0, "Error al momento de convertir, por davor digite un valor que represente un numero como tal"
	}
	if res > 100 {
		return res, "Es mayor a 100"
	} else {
		return res, "Es menor a 100"
	}

}
