package files

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Faydiamond/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabarTabla() {
	texto := ejercicios.Multiplicar()
	archivo, err := os.Create(filename) //pisa archivo
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return //salir de la funcion
	}

	fmt.Fprintln(archivo, texto) //guardo en el archivo, lo que viene en texto
	archivo.Close()
}

func SumaTabla() {
	texto := ejercicios.Multiplicar() + "\n"
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el writestring " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(" Error al leer el archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))
}
