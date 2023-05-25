package funciones

import "fmt"

//clouster
func tabla(valor int) func() int {
	numero := valor
	secueencia := 0

	return func() int {
		secueencia++
		return numero * secueencia
	}
}

//ventaja ofuscar codigo
func LlamarClousure() {
	tabladel := 2
	MiTabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
