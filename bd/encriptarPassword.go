package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/*EncriptarPassword encriptar el password*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //cantidad 2 elevado al costo
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
