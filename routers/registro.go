package routers

import (
	"encoding/json"
	"net/http"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"github.com/luisesteban91/curso_go_api_twiter/models"
)

/*Registro es la funcion para crear en la BD el usuario*/
func Registro(resW http.ResponseWriter, req *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(req.Body).Decode(&t) // asignar el body a t
	if err != nil {
		http.Error(resW, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(resW, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Email) < 6 {
		http.Error(resW, "Debe ser un password de almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(resW, "Ya existe un usuario regisgtrado con este email", 400)
		return
	}

	_, status, err := bd.InsertarUsuario(t)

	if err != nil {
		http.Error(resW, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(resW, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	resW.WriteHeader(http.StatusCreated)
}
