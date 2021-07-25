package main

import "fmt"

func main() {
	fmt.Println("------------------")
	fmt.Println(uno(5))

	fmt.Println("------------------")
	numero, estado := dos(1)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println("------------------")
	fmt.Println(tres(5))

	fmt.Println("parametrosVariables")
	fmt.Println(parametrosVariables(5, 5))
	fmt.Println(parametrosVariables(5, 20))
	fmt.Println(parametrosVariables(5, 20, 2))
	fmt.Println(parametrosVariables(4, 6, 5))
}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func tres(numero int) (z int) {
	z = numero * 2 //por default retornara z ya que esta declarada como retorno
	return
}

func parametrosVariables(numero ...int) int { //funcion tipo de parametros variables
	total := 0
	for _, num := range numero { //se puede cabiar el  "_"  por el "item"
		total = total + num
	}

	return total
}
