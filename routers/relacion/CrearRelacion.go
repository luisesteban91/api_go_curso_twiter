package relacion

import (
	"net/http"

	"github.com/luisesteban91/curso_go_api_twiter/bd/relacion"
	"github.com/luisesteban91/curso_go_api_twiter/models"
	procesarToken "github.com/luisesteban91/curso_go_api_twiter/routers"
)

/*CrearRelacion crear realacion en la bd de usuario con usuario*/
func CrearRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = procesarToken.IDUsuario
	t.UsuarioRelacionID = ID

	status, err := relacion.Create(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al crear relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se a logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
