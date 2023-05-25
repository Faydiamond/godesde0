package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 9, 12, 15, 25}

func MuestraSlice() {

	fmt.Println(tablaSlice)
	porcion := arreglo[3:]  //desde posicion 3 al final
	porcion2 := arreglo[:2] //desde inicio, la ultima no la cuenta
	porcion3 := arreglo[6:8]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20) //cinco elementos y 20 de capacidad
	fmt.Printf("Largo %d , capacidad %d ", len(elementos), cap(elementos))
	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d , capacidad %d ", len(nums), cap(nums))
}
