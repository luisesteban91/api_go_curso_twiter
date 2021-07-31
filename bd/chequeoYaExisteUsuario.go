package bd

import (
	"context"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario verificar si existe el usuario*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar el timeout

	db := MongoCon.Database("twiter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}
	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
