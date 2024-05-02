package DB

const (
	DirecionBaseDatos      = "./DB/inventario.db"
	ExisteProducto         = "SELECT EXISTS(SELECT 1 FROM productos WHERE nombre=?)"
	InsertarEnInventario   = "INSERT INTO productos(codigo, nombre, categoria, precio, stock, proveedor) VALUES(?,?,?,?,?,?);"
	SelectTodosProductos   = "SELECT * FROM productos"
	SelectProductoByID     = "SELECT * FROM productos WHERE id=?;"
	SelectNombreSiExiste   = "SELECT nombre FROM productos WHERE nombre=?;"
	UpdateProduc           = "UPDATE productos SET codigo=?, nombre=?, categoria=?, precio=?, stock=?, proveedor=? WHERE id=?;"
	DeleteProduct          = "DELETE FROM productos WHERE id=?"
	ExisteUsuario          = "SELECT EXISTS(SELECT 1 FROM usuarios WHERE nombre=?)"
	CrearNuevoUsuario      = "INSERT INTO usuarios(nombre, email, contraseña, rol, fecha_creacion, fecha_modificacion, EstadoSesion) VALUES(?,?,?,?,?,?,?);"
	SelectUserByNombre     = "SELECT contraseña FROM usuarios WHERE nombre=?;"
	SelectTodosUser        = "SELECT * FROM usuarios"
	SelectEstadoSecionUser = "SELECT EstadoSesion FROM usuarios where nombre=?"
	SelectRolUser          = "SELECT rol FROM usuarios where nombre=?"
	UpdatetEstadoUser      = "UPDATE usuarios SET EstadoSesion=? WHERE nombre=?;"
	UpdateUser             = "UPDATE usuarios SET nombre=?, email=?, contraseña=?, rol=?, fecha_modificacion=? WHERE id=?;"
	DeleteUser             = "DELETE FROM usuarios WHERE id=?"
)

type ProStru struct {
	Id        int
	Codigo    int
	Nombre    string
	Categoria string
	Precio    float64
	Stock     int
	Proveedor string
}

type Usuarios struct {
	Id                 int
	Nombre             string
	Email              string
	Contraseña         string
	Rol                string // Rol (administrador o general)
	Fecha_creacion     string
	Fecha_modificacion string
	EstadoSesion       string // conectado o Desconectado
}

type Secion struct {
	Id         int
	Idusuario  int
	HoraInicio string
	HoraFin    string
}

var TablasInventario = map[string]string{
	"usuarios": `CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT UNIQUE,
		email TEXT,
		contraseña TEXT,
		rol TEXT,
		fecha_creacion TEXT,
		fecha_modificacion TEXT,
		EstadoSecion VARCHAR(10));`,
	"secion": `CREATE TABLE IF NOT EXISTS siones (
			id INT PRIMARY KEY,
			idusuario INT,
			HoraInicio TEXT,
			HoraFin TEXT,
			FOREIGN KEY (idusuario) REFERENCES Usuarios(id));`,
	"productos": `CREATE TABLE IF NOT EXISTS productos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		codigo INTEGER UNIQUE,
		nombre TEXT UNIQUE,
		categoria TEXT,
		precio REAL,
		stock INTEGER,
		proveedor TEXT);`,
	"entradas": `CREATE TABLE IF NOT EXISTS entradas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fecha TEXT,
		codigo TEXT,
		producto TEXT,
		cantidad INTEGER,
		precio REAL,
		proveedor TEXT,
		factura INTEGER,
		FOREIGN KEY(producto) REFERENCES productos(nombre));`,
	"salidas": `CREATE TABLE IF NOT EXISTS salidas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fecha TEXT,
		codigo TEXT,
		producto TEXT,
		cantidad INTEGER,
		precio REAL,
		cliente TEXT,
		factura INTEGER,
		FOREIGN KEY(producto) REFERENCES productos(nombre));`,
}
