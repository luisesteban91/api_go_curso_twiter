package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCon es el objecto de conexion a la BD*/
var MongoCon = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://dbLecch:1jIcM6SxWVgy9Ys7@cluster0.w4xh2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*
* ConnectBD es la funcion para conectar con la BD
* return client
 */
func ConnectDB() *mongo.Client {
	Client, err := mongo.Connect(context.TODO(), clientOptions) //context.TODO() siginifica sin restricion de timeout, ect

	if err != nil {
		log.Fatal(err.Error())
		return Client
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return Client
	}

	log.Println("Conexion Exitosa con la BD")
	return Client
}

/*
*CheckConnection funcion que valida un ping a la BD
*
 */
func CheckConnection() bool {
	err := MongoCon.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
