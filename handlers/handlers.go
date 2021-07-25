package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT") //obetener variable de entorno

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) //establecer todos los ṕermisos de cors
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
