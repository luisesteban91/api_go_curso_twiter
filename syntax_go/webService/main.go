package main

import "net/http"

func main() {
	http.HandleFunc("/", home)        //crear la ruta
	http.ListenAndServe(":3000", nil) //utilizar el puerto 3000 para recibir las peticiones
}

func home(res http.ResponseWriter, req *http.Request) { //recibir response y request
	http.ServeFile(res, req, "./index.html")
}
