package bd

import (
	"context"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarUsuario es la parada final con la BD para insertar los datos del usuario*/
func InsertarUsuario(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar el timeout

	db := MongoCon.Database("twiter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID) //obtener el object id
	return ObjID.String(), true, nil                   //convertir el object id a string
}
