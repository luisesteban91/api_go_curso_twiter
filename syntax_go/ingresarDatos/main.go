package main

import (
	"bufio"
	"fmt"
	"os"
)

var num1 int
var num2 int
var descripcion string

func main() {
	fmt.Println("Ingresa un numero:")
	fmt.Scanln(&num1)
	fmt.Println("Ingresa un numero:")
	fmt.Scanln(&num2)

	fmt.Println("ingresa una descripcion:")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		descripcion = scanner.Text()
	}
	fmt.Println(num1 + num2)

	fmt.Println(descripcion, num1+num2)
}
