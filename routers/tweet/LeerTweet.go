package tweet

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/luisesteban91/curso_go_api_twiter/bd/tweet"
)

/*Leer ller tweets*/
func Leer(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) //convertir de alfa a numerico
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor", http.StatusBadRequest)
		return
	}

	pag := int64(pagina) //la rutina para paginar requiere en int64
	respuesta, correcto := tweet.GetTweet(ID, pag)

	if correcto == false {
		http.Error(w, "Error al leet los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
