package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ejemploDefer()
	ejemploPanic()
}

func ejemploDefer() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close() //defer siempre se va a ejecutar falle o no falle la ejecucion

	if err != nil {
		fmt.Println("Error al abrir el archivo")
		os.Exit(1)
	}
}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero Panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}
