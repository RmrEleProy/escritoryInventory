package DB

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DBConnection() (db *sql.DB) {
	db, err := sql.Open("sqlite3", DirecionBaseDatos)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

// retorna "existe" si encuentra la tabla  o
// "No existe" si no encuentra la tabla
// y muestra error si ocurre alguno
func VerificaTablas(tabla string) (string, error) {
	db := DBConnection()
	defer db.Close()

	sqlStmt := `SELECT name FROM sqlite_master WHERE type='table' AND name=?;`
	err := db.QueryRow(sqlStmt, tabla).Scan(&tabla)
	if err != nil {
		if err == sql.ErrNoRows {
			return "No Existe", nil
		} else {
			return "Ha ocurrido un error", err
		}
	} else {
		return "Existe", nil
	}
}

func CrearTablas(sqlStm string) {
	db := DBConnection()
	defer db.Close()
	_, err := db.Exec(sqlStm)

	if err != nil {
		fmt.Println(err)
	}
}

// esta funcion Crea las tablas en la base de datos inventario
// las tablas: Usuarios, productos, entradas, salidas
func FuncionInicial() {
	for nombre, sqlStm := range TablasInventario {
		hayt, err := VerificaTablas(nombre)
		if err != nil {
			fmt.Println(err)
		}
		if hayt == "No Existe" {
			CrearTablas(sqlStm)
		}
	}
}

// verifica si ya existe un registro en la base de datos.
func RegistroExiste(nombre, sqlstm string) bool {
	db := DBConnection()
	var existe bool
	err := db.QueryRow(sqlstm, nombre).Scan(&existe)
	if err != nil {
		fmt.Println(err)
	}
	return existe
}

// inserta en la base de datos los nuevos productos
func InsertarNewRegistro(p ProStru) {
	if !RegistroExiste(p.Nombre, ExisteProducto) {
		db := DBConnection()
		defer db.Close()

		insForm, err := db.Prepare(InsertarEnInventario)
		if err != nil {
			panic(err.Error())
		}
		defer insForm.Close()
		_, err = insForm.Exec(p.Codigo, p.Nombre, p.Categoria, p.Precio, p.Stock, p.Proveedor)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("ya existe este producto")
	}

}

// funcion para ver todos los productos
func VerTodosLosRegistros() []ProStru {
	db := DBConnection()
	defer db.Close()
	seldb, err := db.Query(SelectTodosProductos)
	if err != nil {
		fmt.Println(err.Error())
	}

	prod := ProStru{}
	prodVector := []ProStru{}
	for seldb.Next() {
		var id, codigo, stock int
		var precio float64
		var nombre, categoria, proveedor string
		err = seldb.Scan(&id, &codigo, &nombre, &categoria, &precio, &stock, &proveedor)
		if err != nil {
			panic(err.Error())
		}
		prod.Id = id
		prod.Codigo = codigo
		prod.Nombre = nombre
		prod.Categoria = categoria
		prod.Precio = precio
		prod.Stock = stock
		prod.Proveedor = proveedor

		prodVector = append(prodVector, prod)
	}
	return prodVector
}

// funcion que muestra un producto asociado a un id selecionado
func VerOneRegistro(id int) (p ProStru) {
	db := DBConnection()
	defer db.Close()
	selDB, err := db.Query(SelectProductoByID, id)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, codigo, stock int
		var precio float64
		var nombre, categoria, proveedor string
		err = selDB.Scan(&id, &codigo, &nombre, &categoria, &precio, &stock, &proveedor)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Codigo = codigo
		p.Nombre = nombre
		p.Categoria = categoria
		p.Precio = precio
		p.Stock = stock
		p.Proveedor = proveedor
	}
	return p
}

// funcion que actualiza un producto con un ID seleccionado
func EditarProducto(p ProStru) {
	db := DBConnection()
	defer db.Close()
	insForm, err := db.Prepare(UpdateProduc)
	if err != nil {
		panic(err.Error())
	}
	_, err = insForm.Exec(p.Codigo, p.Nombre, p.Categoria, p.Precio, p.Stock, p.Proveedor, p.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Elimina un registro Con el ID seleccionado
func BorrarOldRegistro(ID int) {
	bd := DBConnection()
	defer bd.Close()

	delForm, err := bd.Prepare(DeleteProduct)
	if err != nil {
		panic(err.Error())
	}
	_, err = delForm.Exec(ID)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//**#######################USER######################

// crea nuevos usuarios
func InsertarNewUser(u Usuarios) string {
	if !RegistroExiste(u.Nombre, ExisteUsuario) {
		db := DBConnection()
		defer db.Close()

		insForm, err := db.Prepare(CrearNuevoUsuario)
		if err != nil {
			return err.Error()
		}
		defer insForm.Close()
		_, err = insForm.Exec(u.Nombre, u.Email, u.Contraseña, u.Rol, u.Fecha_creacion, u.Fecha_modificacion, u.EstadoSesion)
		if err != nil {
			return err.Error()
		}
		return "se agrego el usuario: " + u.Nombre
	} else {
		return "la persona que intenta agregar ya ha sido registrada"
	}
}

// retorna el hass de la contraseña encriptada generada con bycript, del usuario registrado
func Buscarhass(nombre string) (contraseña string) {
	db := DBConnection()
	defer db.Close()

	selDB, err := db.Query(SelectUserByNombre, nombre)
	if err != nil {
		fmt.Println(err.Error())
	}

	for selDB.Next() {
		err = selDB.Scan(&contraseña)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return contraseña
}

// retorna el estado secion de un usuario
func EstadoSesionUsuario(nombre string) (estadoSesion string) {
	db := DBConnection()
	defer db.Close()

	Estadodb, err := db.Query(SelectEstadoSecionUser, nombre)
	if err != nil {
		fmt.Println(err)
	}

	for Estadodb.Next() {
		err = Estadodb.Scan(&estadoSesion)
		if err != nil {
			fmt.Println(err)
		}
	}
	return estadoSesion
}

// Retorna el Rol de un usuario Administrador O General
// requiere el nombre del usuario
func RolUser(nombre string) (rol string) {
	db := DBConnection()
	defer db.Close()

	rolDBuser, err := db.Query(SelectRolUser, nombre)
	if err != nil {
		fmt.Println(err)
	}

	for rolDBuser.Next() {
		err = rolDBuser.Scan(&rol)
		if err != nil {
			fmt.Println(err)
		}
	}
	return rol
}

// actualiza el estado del usuario correspondiente
func UpdateSesionUser(estado, nombre string) {
	db := DBConnection()
	defer db.Close()

	sesionUserdb, err := db.Prepare(UpdatetEstadoUser)
	if err != nil {
		fmt.Println(err)
	}
	_, err = sesionUserdb.Exec(estado, nombre)
	if err != nil {
		fmt.Println(err)
	}
}

// funcion que actualiza un usuario con un ID seleccionado
func EditarUsuario(user Usuarios) {
	db := DBConnection()
	defer db.Close()
	insForm, err := db.Prepare(UpdateUser)
	if err != nil {
		panic(err.Error())
	}
	_, err = insForm.Exec(user.Nombre, user.Email, user.Contraseña, user.Rol, user.Fecha_modificacion, user.Id)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Elimina un usuario asociado a un ID proporcionado
func DeleteUserbyId(Id int) {
	db := DBConnection()
	defer db.Close()

	stmdel, err := db.Prepare(DeleteUser)
	if err != nil {
		fmt.Println(err)
	}

	_, err = stmdel.Exec(Id)
	if err != nil {
		fmt.Println(err)
	}
}

// Retorna todos los usuarios almacenados en la base de datos
func VerTodosLosUsuarios() []Usuarios {
	user := Usuarios{}
	uservector := []Usuarios{}

	db := DBConnection()
	defer db.Close()

	seldb, err := db.Query(SelectTodosUser)
	if err != nil {
		fmt.Println(err)
	}

	for seldb.Next() {
		var id int
		var nombre, email, contraseña, rol, fecha_creacion, fecha_modificacion, EstadoSecion string
		err = seldb.Scan(&id, &nombre, &email, &contraseña, &rol, &fecha_creacion, &fecha_modificacion, &EstadoSecion)
		if err != nil {
			fmt.Println(err)
		}
		user.Id = id
		user.Nombre = nombre
		user.Email = email
		user.Contraseña = contraseña
		user.Rol = rol
		user.Fecha_creacion = fecha_creacion
		user.Fecha_modificacion = fecha_modificacion
		user.EstadoSesion = EstadoSecion

		uservector = append(uservector, user)
	}

	return uservector
}
