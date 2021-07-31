package main

import "fmt"

func main() {
	tablaDel := 2
	MiTabla := Tabla(tablaDel)

	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}

func Tabla(valor int) func() int { //funcion tabla devuelve un dato de tipo funcion
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
