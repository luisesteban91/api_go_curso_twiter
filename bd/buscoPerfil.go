package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca un perfil en la bd*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar el timeout
	db := MongoCon.Database("twiter")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	ObjectID, _ := primitive.ObjectIDFromHex(ID) //hex hexadecimal

	condicion := bson.M{
		"_id": ObjectID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
