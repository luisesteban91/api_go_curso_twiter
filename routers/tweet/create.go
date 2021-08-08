package tweet

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/bd/tweet"
	"github.com/luisesteban91/curso_go_api_twiter/models"
	procesarToken "github.com/luisesteban91/curso_go_api_twiter/routers"
)

/*Create crea un tweet en la BD*/
func Create(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.CrearTweet{
		UserId:  procesarToken.IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := tweet.Create(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
