package main

import (
	"fmt"
	"time"

	us "estructuras/user" //en el ejemplo de la clase seria 'ejer09' por 'estructuras'
)

type pepe struct {
	us.Usuario //herencia
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Luis Esteban", time.Now(), true)

	fmt.Println(u.Usuario)
}
