package tweet

import (
	"context"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"github.com/luisesteban91/curso_go_api_twiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Create crear un tweet en la BD*/
func Create(t models.CrearTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar el timeout
	db := bd.MongoCon.Database("twiter")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserId,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro) //crear tweet

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
