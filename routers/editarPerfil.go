package routers

import (
	"encoding/json"
	"net/http"

	"github.com/luisesteban91/curso_go_api_twiter/bd/usuario"
	"github.com/luisesteban91/curso_go_api_twiter/models"
)

/* ModificarPerfil modifica el perfil del usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var us models.Usuario

	err := json.NewDecoder(r.Body).Decode(&us) //asignar el usaurio que viene en el request en la variable us

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = usuario.EditarUsuario(us, IDUsuario) //IDUsuario variable glogal que se crea al decodificar el token
	if err != nil {
		http.Error(w, "Ocurrio unn error al modificar le registro."+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se a logrado modificar el registro de usaurio ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
