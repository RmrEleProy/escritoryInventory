package Controles

import (
	"Inventario/DB"
	"regexp"
	"time"
	"unicode"
)

// Actualiza el estado de secion del usuario
func RevisarUser(user string) {
	// switch DB.EstadoSesionUsuario(user) {
	// case "Activo":
	// 	DB.UpdateSesionUser("Desconectado", user)
	// case "Desconectado":
	// 	DB.UpdateSesionUser("Activo", user)
	// case " ":
	// 	DB.UpdateSesionUser("Desconectado", user)
	// default:
	// 	DB.UpdateSesionUser("Desconectado", user)
	// }
}

// Valida los campos para el ingreso de los nuevos usuarios
func ValidarCamposNewUser(user DB.Usuarios) string {
	// rFecha, _ := regexp.Compile(`^\d{4}-\d{2}-\d{2}$`)
	fechaLayout := "2/01/2006"
	rCorreo := regexp.MustCompile(`(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))`)

	if user.Nombre == "" {
		return "El nombre no puede esta vacio"
	}

	if len(user.Nombre) < 4 {
		return "el nombre debe tener almenos 4 caractenes"
	}

	if !rCorreo.MatchString((user.Email)) {
		return "Correo no valido"
	}
	if validarContraseña(user.Contraseña) != "ok" {
		return validarContraseña(user.Contraseña)
	}

	if user.Rol != "Administrador" && user.Rol != "General" {
		return ("Rol debe ser 'Administrador' o 'General'")
	}

	_, err := time.Parse(fechaLayout, user.Fecha_creacion)
	if err != nil {
		return "Fecha no valida -- el formato de la fecha debe ser 2/01/2006"
	}

	return "ok"
}

// todo verificar esta funcion
func validarContraseña(contraseña string) string {
	numeros := 0
	mayus := 0
	minus := 0
	special := 0
	spacio := 0
	for _, c := range contraseña {
		switch {
		case unicode.IsNumber(c):
			numeros++
		case unicode.IsUpper(c):
			mayus++
		case unicode.IsLower(c):
			minus++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special++
		case unicode.IsSpace(c):
			spacio++
		}
	}
	if len(contraseña) <= 7 {
		return "La contraseña debe tener minimo 8 caracteres"
	}
	if mayus == 0 {
		return "Debe tener almenos una mayuscula"
	}
	if minus == 0 {
		return "Debe tener almenos una minuscula"
	}
	if numeros == 0 {
		return "Debe tener almenos un numero"
	}
	if special == 0 {
		return "Debe tener almenos un caracter expecial"
	}
	if spacio != 0 {
		return "No debe contener espacio"
	}
	return "ok"
}
