package main

import (
	"log"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"github.com/luisesteban91/curso_go_api_twiter/handlers"
)

func main() {
	if bd.CheckConnection() == false {
		log.Fatal("sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
