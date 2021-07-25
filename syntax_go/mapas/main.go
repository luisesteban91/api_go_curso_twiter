package main

import "fmt"

func main() {
	capitales := make(map[string]string, 5)

	fmt.Println(capitales)
	capitales["Mexico"] = "DF"
	capitales["Argentina"] = "Buenos Aires"
	fmt.Println(capitales)

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}
	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 55
	delete(campeonato, "Real Madrid") //borrar
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato { //Recorrer un mapa
		fmt.Printf("El equipo %s, tiene un putaje de : %d \n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
