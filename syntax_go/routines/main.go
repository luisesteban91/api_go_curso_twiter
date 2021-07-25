package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Esteban") //go va a ejecutar la funcion pero de manera asicrona
	fmt.Println("estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "") //convetir letras en un vector

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
