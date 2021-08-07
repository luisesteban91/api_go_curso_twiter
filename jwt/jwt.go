package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/luisesteban91/curso_go_api_twiter/models"
)

func GetJWT(t models.Usuario) (string, error) {
	miClave := []byte("MastersDelDesarrollo")
	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fechaNacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"sitioweb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //algoritmo para encriptar HS256
	tokenStr, err := token.SignedString(miClave)                //firmar
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
