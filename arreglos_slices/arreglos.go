package arreglos_slices

import "fmt"

//arreglos o vectores

var tabla [10]int = [10]int{5, 10, 15, 20, 25}
var matriz [20][30]int

func MuetraArreglos() {

	tabla2 := [4]string{"Hola", "Vamos a ", "intentarlo", " de nuevo."}

	fmt.Println(tabla)

	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matriz[5][25] = 15
	fmt.Println(matriz)
}
