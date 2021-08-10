package relacion

import (
	"net/http"

	"github.com/luisesteban91/curso_go_api_twiter/bd/relacion"
	"github.com/luisesteban91/curso_go_api_twiter/models"
	procesarToken "github.com/luisesteban91/curso_go_api_twiter/routers"
)

func BorrarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = procesarToken.IDUsuario
	t.UsuarioRelacionID = ID

	status, err := relacion.Delete(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se a logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
