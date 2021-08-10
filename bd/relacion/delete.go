package relacion

import (
	"context"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"github.com/luisesteban91/curso_go_api_twiter/models"
)

/*Delete borrar una relacion en la BD*/
func Delete(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar el timeout
	db := bd.MongoCon.Database("twiter")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
