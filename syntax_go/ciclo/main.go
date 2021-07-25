package main

import "fmt"

func main() {
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for x := 0; x <= 10; x++ {
	// 	fmt.Println(x)
	// }

	// num := 0

	// for {
	// 	fmt.Println("continuo")
	// 	fmt.Println("ingrese el numero secreto")
	// 	fmt.Scanln(&num)
	// 	if num == 100 {
	// 		break
	// 	}
	// }

	// var y = 0
	// for y < 10 {
	// 	fmt.Printf("\n Valor de i: %d", y)
	// 	if y == 5 {
	// 		fmt.Printf("multiplicamos por 2 \n")
	// 		y = 1 * 2
	// 		continue // se brinca al siguiente bloque del for
	// 	}
	// 	fmt.Printf("       Pasó por aqui \n")
	// 	y++
	// }

	var p int = 0

RUTINA:
	fmt.Println("entro a RUTINA")
	for p < 10 {
		if p == 4 {
			p = p + 2
			fmt.Println("voy a RUTINA sumando 2 a p")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", p)
		p++
	}
}
