package Controles

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Encriptar(password string) string {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(hashpassword)
}

// compara la contraseña encriptada con la contraseña proporcionada
// retorna TRUE o FALSE
func VerificaPasword(hashpassrd, currpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassrd), []byte(currpassword))
	return err == nil
}
