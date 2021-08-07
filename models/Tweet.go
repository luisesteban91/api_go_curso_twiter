package models

import "time"

/*CrearTweet modelo de crear un tweet*/
type CrearTweet struct {
	UserId  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}

/*Tweet modelo que captura el body, el mensaje que nos llega*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
