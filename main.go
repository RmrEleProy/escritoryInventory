package main

import (
	"Inventario/Controles"
	"Inventario/DB"
	"Inventario/Vista"

	"fmt"
)

func main() {
	DB.FuncionInicial()

	user := DB.Usuarios{
		Nombre:             "Admin",
		Email:              "admin@inventary.com",
		Contraseña:         Controles.Encriptar("Adminpass"),
		Rol:                "Administrador",
		Fecha_creacion:     "07/03/2024",
		Fecha_modificacion: "07/03/2024",
		EstadoSesion:       "activo"}
	fmt.Println(DB.InsertarNewUser(user))

	Vista.Inicio()
}

func PreguntaTecla() int {
	var numero int
	fmt.Println("selecciona un registro: ")
	fmt.Scanf("%d", &numero)
	return numero

}
