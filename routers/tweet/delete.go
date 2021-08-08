package tweet

import (
	"net/http"

	"github.com/luisesteban91/curso_go_api_twiter/bd/tweet"
	procesarToken "github.com/luisesteban91/curso_go_api_twiter/routers"
)

/*Delete elimininar tweet*/
func Delete(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := tweet.Delete(ID, procesarToken.IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un erro al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
