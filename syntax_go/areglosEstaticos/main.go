package main

import "fmt"

var vector [10]int

func main() {
	tabla := [10]int{5, 6, 7, 8, 9, 12, 12, 13, 23}

	tabla[0] = 1
	tabla[5] = 15

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	var matriz [5][7]int
	matriz[2][3] = 1

	fmt.Println(matriz)
}
