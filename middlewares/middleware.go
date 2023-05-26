package middlewares

import "fmt"

// son funciones que se ejecutan antes de determinada funciion
// requisito lo que reciben es del mismo tipo de dato de lo que retorna

func sumar(n1, n2 int) int {
	return n1 + n2
}

func restar(n1, n2 int) int {
	return n1 - n2
}

func multiplicar(n1, n2 int) int {
	return n1 * n2
}

//funcion que ejecuta otras funciones
func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMid(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMid(restar)(5, 3)
	fmt.Println(result)
	result = operacionesMid(multiplicar)(2, 3)
	fmt.Println(result)
}

func operacionesMid(f func(int, int) int) func(int, int) int { //solo va el tipado de datos
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
