package tweet

import (
	"context"
	"log"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"github.com/luisesteban91/curso_go_api_twiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options" //filtrar en la bd
)

/* GetTweet leer los comentarios de un perfil*/
func GetTweet(ID string, paginate int64) ([]*models.DevuelvoTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar el timeout
	db := bd.MongoCon.Database("twiter")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweet

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)                               //cuantos documentos me va a atraer
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //ordenar por fecha
	opciones.SetSkip((paginate - 1) * 20)               //ir paginando

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweet
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true

}
