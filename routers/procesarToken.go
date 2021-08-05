package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/luisesteban91/curso_go_api_twiter/bd"
	"github.com/luisesteban91/curso_go_api_twiter/models"
)

/*Email valor del Email usado en todos los endpoints*/
var Email string

/*IDUsuario valor del IDUsuario usado en todos los endpoints*/
var IDUsuario string

/*ProcesarToken estraer valores del token*/
func ProcesarToken(miToken string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterdelDesarrollo")
	claims := &models.Claim{}

	splitToken := strings.Split(miToken, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token no valido") //errors no permite signos(?¡)
	}

	miToken = strings.TrimSpace(splitToken[1]) //obtener el puro token

	tkn, err := jwt.ParseWithClaims(miToken, claims, func(token *jwt.Token) (interface{}, error) { //ParseWithClaims parsear el token lo mapea dentro de claims y checar si es valido
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Invalido")
	}

	return claims, false, string(""), err
}
