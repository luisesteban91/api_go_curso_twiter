package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leerARchivo()
	leerARchivo2()
	crearArchivo()
	crearArchivo2()
}

func leerARchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt") //este se encarga de abrir y cerrar el archivo
	if err != nil {
		fmt.Println("Hubo un error")
	}

	fmt.Println(string(archivo))
}

func leerARchivo2() {
	archivo, err := os.Open("./archivo.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Printf("Linea > " + registro + "\n")
	}

	archivo.Close()
}

func crearArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	archivo.Close()
}

func crearArchivo2() {
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\n Esta es una segunda linea") == false { //agrear texto a un archivo
		fmt.Println("Error en la 2da linea")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	_, err = arch.WriteString(texto) //el _, es cuando no nos interasa el primer valor que devuelve una una funcion
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	return true
}
