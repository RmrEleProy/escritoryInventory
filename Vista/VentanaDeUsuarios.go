package Vista

import (
	"Inventario/DB"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Usuarios generales
func VentanaUsuarios(a fyne.App, rol, nombre string) {
	titulo := "Ventana de opciones de usuario: " + rol
	w := GenerarVentana(a, titulo, 500, 450, true)
	// colocar en el titulo de la ventana que aparezca el nombre o si es general
	UserLogin := "sesion iniciada de: " + nombre
	LabelTitulo := widget.NewLabel(UserLogin)
	ContenedorTitulo := container.NewHBox(layout.NewSpacer(), LabelTitulo, layout.NewSpacer())

	BtnVenta := widget.NewButton("Nueva venta", func() {
		Login(a)
	})
	BtnInventario := widget.NewButton("Ver Inventario", func() {
		VentanaInventarios(a, DB.RolUser(rol))
	})
	BtnCompras := widget.NewButton("Nueva compra", func() {
	})
	ContenedorBtns := container.NewHBox(BtnInventario, BtnVenta, BtnCompras)

	fotoPerfil := canvas.NewImageFromFile("static/IMG/Designer.png")
	fotoPerfil.FillMode = canvas.ImageFillOriginal

	contenido := container.NewVBox(ContenedorTitulo, ContenedorBtns, fotoPerfil)

	menu1 := fyne.NewMenu("File",
		fyne.NewMenuItem("Cerrar sesion", func() {
			Login(a)
			w.Hide()
		}),
		fyne.NewMenuItem("Opción 1.2", func() {}),
	)

	menu2 := fyne.NewMenu("Submenu 2",
		fyne.NewMenuItem("Opción 2.1", func() {}),
		fyne.NewMenuItem("Opción 2.2", func() {}),
	)

	mainMenu := fyne.NewMainMenu(menu1, menu2)

	w.SetMainMenu(mainMenu)

	w.SetContent(contenido)
	w.Show()
	w.SetOnClosed(func() { a.Quit() })
}

// todo terminar de arreglar esta ventana
func VentanaAdmin(a fyne.App, rol, nombre string) {
	w := GenerarVentana(a, "Inventario Admin", 400, 350, true)

	fondo := canvas.NewImageFromFile("static/IMG/ADmin.png")
	fondo.FillMode = canvas.ImageFillOriginal

	menu1 := fyne.NewMenu("Usuarios",
		fyne.NewMenuItem("Agregar Usuario", func() {
			NewUser(a)
		}),
		fyne.NewMenuItem("Ver Usuarios", func() { ListarUsuarios(a) }),
	)

	menu2 := fyne.NewMenu("Inventario",
		fyne.NewMenuItem("agregar un nuevo articulo", func() {

		}),
		fyne.NewMenuItem("Ver articulos de inventario", func() {
			VentanaInventarios(a, DB.RolUser(rol))
		}),
	)

	mainmenu := fyne.NewMainMenu(menu1, menu2)
	w.SetMainMenu(mainmenu)
	w.Show()
	w.SetOnClosed(func() {
		a.Quit()
	})
}
