package mapas

import (
	"fmt"
)

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Argentina"] = "Buenos Aires"
	paises["México"] = "D.F."
	paises["Uruguay"] = "Montevideo"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"river":         42,
		"boca":          39,
		"independiente": 5,
		"huracán":       21,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene %d puntos \n", equipo, puntaje)
	}

	delete(campeonato, "independiente")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["independiente"]
	fmt.Printf("Esquipo existe: %t, puntaje: %d", existe, puntaje)
}
