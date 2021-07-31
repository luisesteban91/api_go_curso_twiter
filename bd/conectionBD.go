package bd

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCon es el objecto de conexion a la BD*/
var MongoCon = ConnectDB()

/*
* ConnectBD es la funcion para conectar con la BD
* return client
 */
func ConnectDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GO_PASSWORD_BD := os.Getenv("GO_PASSWORD_BD") //obetener variable de entorno

	var clientOptions = options.Client().ApplyURI("mongodb+srv://dbLecch:" + GO_PASSWORD_BD + "@cluster0.w4xh2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

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
