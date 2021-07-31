package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Ususario es el modelo de usuario de la BD MongoDB*/
type Usuario struct {
	// dato input, dato output
	// bson:"_id,  omitempty" json:"id"
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"` //omitir si esta vacio el id
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"` //jamas se debe pasar el password por el navegador
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
}
