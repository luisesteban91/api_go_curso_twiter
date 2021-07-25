package main

import (
	"fmt"
	"strconv"
)

var number int
var text string
var status bool = true

func main() {
	//estas variables las trata de manera local solo en la funcion
	// el := significa crear una nueva variable
	number2, number3, number4, text, status := 1, 5, 67, "Este es mi texto", false

	var currency float32 = 0
	//number2 = currency// esto mandara error por que currency(moneda) es float y number2 se declaro como int

	number2 = int(currency) //convertir currency(moneda) en int

	text = strconv.Itoa(int(currency)) //convertir float a string

	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4)
	fmt.Println(text)
	fmt.Println(status)

	showStatus()
}

func showStatus() {
	fmt.Println(status)
}
