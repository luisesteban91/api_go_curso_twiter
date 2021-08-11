package tweet

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/luisesteban91/curso_go_api_twiter/bd/tweet"
	procesarToken "github.com/luisesteban91/curso_go_api_twiter/routers"
)

/*Leer tweets de nuestros seguidores*/
func LeerTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("pagina")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entreo mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := tweet.ReadTweetsFollows(procesarToken.IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
