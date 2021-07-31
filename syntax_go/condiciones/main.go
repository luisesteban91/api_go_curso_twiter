package main

import "fmt"

var status bool

func main() {
	status = true
	if status {
		fmt.Println(status)
	} else {
		fmt.Println("esto es false")
	}
}
