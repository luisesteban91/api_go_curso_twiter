package middleware

import (
	"net/http"

	"github.com/luisesteban91/curso_go_api_twiter/routers"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesarToken(r.Header.Get("Autorization"))
		if err != nil {
			http.Error(rw, "Token no valido ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
