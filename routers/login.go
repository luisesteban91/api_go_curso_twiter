package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"github.com/luisesteban91/curso_go_api_twiter/jwt"
	"github.com/luisesteban91/curso_go_api_twiter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	usuario, login := bd.Login(t.Email, t.Password)
	if login == false {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	//crear token jwt
	jwtKey, err := jwt.GetJWT(usuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token correspondiente "+err.Error(), 400)
		return
	}

	response := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	//generar cookie desde backend
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
