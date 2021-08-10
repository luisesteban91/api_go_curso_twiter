package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luisesteban91/curso_go_api_twiter/middleware"
	"github.com/luisesteban91/curso_go_api_twiter/routers"
	"github.com/luisesteban91/curso_go_api_twiter/routers/avatar"
	"github.com/luisesteban91/curso_go_api_twiter/routers/banner"
	"github.com/luisesteban91/curso_go_api_twiter/routers/relacion"
	"github.com/luisesteban91/curso_go_api_twiter/routers/tweet"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/miperfil", middleware.ChequeoBD(middleware.ValidateJWT(routers.MiPerfil))).Methods("GET")
	router.HandleFunc("/editarPerfil", middleware.ChequeoBD(middleware.ValidateJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.ChequeoBD(middleware.ValidateJWT(tweet.Create))).Methods("POST")
	router.HandleFunc("/readTweets", middleware.ChequeoBD(middleware.ValidateJWT(tweet.Leer))).Methods("GET")
	router.HandleFunc("/deleteTweet", middleware.ChequeoBD(middleware.ValidateJWT(tweet.Delete))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middleware.ChequeoBD(middleware.ValidateJWT(avatar.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middleware.ChequeoBD(avatar.UploadAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middleware.ChequeoBD(middleware.ValidateJWT(banner.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middleware.ChequeoBD(banner.UploadBanner)).Methods("GET")

	//Relacion
	router.HandleFunc("/crearRelacion", middleware.ChequeoBD(middleware.ValidateJWT(relacion.CrearRelacion))).Methods("POST")
	router.HandleFunc("/deleteRelacion", middleware.ChequeoBD(middleware.ValidateJWT(relacion.BorrarRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middleware.ChequeoBD(middleware.ValidateJWT(relacion.ConsultarRelacion))).Methods("GET")

	PORT := os.Getenv("PORT") //obetener variable de entorno

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) //establecer todos los ṕermisos de cors
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
