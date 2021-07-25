package main

import "fmt"

var calculo func(int, int) int = func(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma 5 + 7 = %d \n", calculo(5, 7))

	//resta
	calculo = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n", calculo(6, 4))

	//divicion
	calculo = func(i1, i2 int) int {
		return i1 / i2
	}
	fmt.Printf("dividimos 6 - 4 = %d \n", calculo(6, 4))

	operaciones()
}

func operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}
