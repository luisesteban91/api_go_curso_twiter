package main

import "fmt"

func main() {
	matriz := []int{2, 3, 4, 5}
	fmt.Println(matriz)

	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]  //toma una porcion del vector de la posición 3 asta la ultima posición
	porcion2 := elementos[:3] //toma una porcion del vector de la posición 0 asta la tercera posición

	fmt.Println(porcion)
	fmt.Println(porcion2)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0) //se ve redimencionando la capacidad en multiplos de 8
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
