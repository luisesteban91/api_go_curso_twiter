package tweet

import (
	"context"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Delete borrar un tweeter por id*/
func Delete(ID string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar el timeout
	db := bd.MongoCon.Database("twiter")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID) //convertir el ID en OnjectId

	condicion := bson.M{
		"_id":    objID,
		"userid": UserId,
	}

	_, err := col.DeleteOne(ctx, condicion)

	return err
}
