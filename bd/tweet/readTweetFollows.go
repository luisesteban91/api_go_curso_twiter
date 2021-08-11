package tweet

import (
	"context"
	"time"

	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"github.com/luisesteban91/curso_go_api_twiter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadTweetsFollows lee los tweets de mis seguidores*/
func ReadTweetsFollows(ID string, pag int) ([]models.GetTweetsFollows, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() //cancelar el timeout

	db := bd.MongoCon.Database("twiter")
	col := db.Collection("relacion")

	skip := (pag - 1) * 20
	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{
		"$match": bson.M{"usuarioid": ID},
	}) //match buscar relacion
	condiciones = append(condiciones, bson.M{
		//crear relacion en mongodb
		"$lookup": bson.M{ //lookap unir 2 tablas
			"from":         "tweet",
			"localField":   "usuariorelacionid", //localfield campo que se va a unir
			"foreignField": "userid",
			"as":           "tweet", //as es el alias(nueva tabla)
		},
	})

	condiciones = append(condiciones, bson.M{
		"$unwind": "$tweet", //unwind que todos los documentos vengan exactamente iguales
	})
	condiciones = append(condiciones, bson.M{
		"$sort": bson.M{"tweet.fecha": -1}, // ordernar por fecha de forma desc
	})
	condiciones = append(condiciones, bson.M{
		"$skip": skip, //cuantos registros debe sortear
	})
	condiciones = append(condiciones, bson.M{
		"$limit": 20,
	})

	cursor, err := col.Aggregate(ctx, condiciones) //Aggregate fraemword de mongo
	var result []models.GetTweetsFollows
	err = cursor.All(ctx, &result) //ejecutar el cursor(armar el documento)
	if err != nil {
		return result, false
	}
	return result, true
}
