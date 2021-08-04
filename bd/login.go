package bd

import (
	"github.com/luisesteban91/curso_go_api_twiter/models"
	"golang.org/x/crypto/bcrypt"
)

/*Login*/
func Login(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBd := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBd, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
