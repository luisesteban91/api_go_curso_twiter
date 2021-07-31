package middleware

import (
	"net/http"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
)

/*ChequeoBD es el middlere que me permite conocer el estado de la BD */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(resW http.ResponseWriter, req *http.Request) {
		if bd.CheckConnection() == false {
			http.Error(resW, "Conexion perdida con la Base de Datos", 500)
			return //Finaliza todo en proceso
		}
		next.ServeHTTP(resW, req)
	}
}
