package main

import (
	e "github.com/Faydiamond/godesde0/ejer_interfaces"
	m "github.com/Faydiamond/godesde0/modelos"
)

func main() {
	/*
		variables.MostrarEnteros()
		fmt.Println("////////////////")
		variables.RestoVariables()
		fmt.Println("////////////////")
		estado, resultado := variables.ConviertoaTexto(2500)
		fmt.Println("Estado: ", estado)
		fmt.Println("Resultado: ", resultado)

		if os := runtime.GOOS; os == "windows" {
			fmt.Println("Su sistema operativo es  Windows")
		}

		switch oss := runtime.GOOS; oss {
		//case "windows":
		//	fmt.Println("Es un Windows")
		case "ubuntu":
			fmt.Println("Es un ubuntu")
		default:
			fmt.Printf("%s  ", oss)
		}
		fmt.Println("////////////////")

		valor, res := ejercicios.Conversiones("XXX")
		fmt.Println("El valor es: ", valor)
		fmt.Printf("El resultado  es: %s ", res)

		iteraciones.Iterar()
	*/
	//files.GrabarTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClousure()
	//funciones.Exponencia(5)
	//arreglos_slices.MuetraArreglos()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AgregarUsuario()
	Mario := new(m.Hombre)
	e.HumanosRespirando(Mario)

	Daniela := new(m.Mujer)
	e.HumanosRespirando(Daniela)
}

//cluster
