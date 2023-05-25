package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["colombia"] = "Bogota D.C"
	paises["Argentina"] = "Buenos Aires"
	paises["Mexico"] = "Mexico D.F"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"barcelona":          40,
		"real madrid":        39,
		"boca juniors":       25,
		"america de mexico ": 23,
		"almeria":            13,
	}

	for equipo, puntos := range campeonato {
		fmt.Println(equipo, puntos)
	}

	delete(campeonato, "almeria")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["chelsea"]

	fmt.Printf(" El puntaje es %d  , el equipo existe %t ", puntaje, existe)
}
