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
	// prod := DB.ProStru{
	// 	Codigo: 2321,
	// 	Nombre:    "naranjas",
	// 	Categoria: "frutas",
	// 	Precio:    20,
	// 	Stock:     10,
	// 	Proveedor: "fruvasa"}
	// DB.InsertarNewRegistro(prod)

	// prod = DB.ProStru{
	// 		Codigo: 1546,
	// 		Nombre:    "peras",
	// 		Categoria: "frutas",
	// 		Precio:    10,
	// 		Stock:     10,
	// 		Proveedor: "fruvasa"}
	// 	DB.InsertarNewRegistro(prod)

	// DB.EditarOldRegistro(prod)

	// fmt.Println(DB.VerTodosLosRegistros())

	Vista.Inicio()
	// fmt.Println("has seleciionado el registro", PreguntaTecla())
	// DB.BorrarOldRegistro(2)
}

func PreguntaTecla() int {
	var numero int
	fmt.Println("selecciona un registro: ")
	fmt.Scanf("%d", &numero)
	return numero

}
