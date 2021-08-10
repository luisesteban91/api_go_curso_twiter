package relacion

import (
	"encoding/json"
	"net/http"

	"github.com/luisesteban91/curso_go_api_twiter/bd/relacion"
	"github.com/luisesteban91/curso_go_api_twiter/models"
	procesarToken "github.com/luisesteban91/curso_go_api_twiter/routers"
)

/*ConsultarRelacion ver la relacion entre 2 usuarios*/
func ConsultarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = procesarToken.IDUsuario
	t.UsuarioRelacionID = ID

	var respuesta models.RespuspuestaConsultaRelacion

	status, err := relacion.ListByUser(t)
	if err != nil || status == false {
		respuesta.Status = false
	} else {
		respuesta.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
