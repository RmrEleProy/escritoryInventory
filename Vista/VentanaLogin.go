package Vista

import (
	"time"

	"Inventario/DB"
	"Inventario/Controles"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Ventana de Login debes enviar una vatiable tipo fyne.app
func Login(a fyne.App) {

	w := GenerarVentana(a, "Inventario general", 400, 350, true)

	clock := widget.NewLabel("")
	updateTime(clock)
	contenedorClock := container.NewHBox(layout.NewSpacer(), clock)

	lebelTitulo := widget.NewLabel(" Login ")
	contenedorTitulo := container.NewHBox(layout.NewSpacer(), lebelTitulo, layout.NewSpacer())

	labelUser := widget.NewLabel("Usuario")
	TextUser := widget.NewEntry()
	ContenedorUser := container.NewVBox(layout.NewSpacer(), labelUser, TextUser, layout.NewSpacer())

	labelPass := widget.NewLabel("Contraseña")
	TextPass := widget.NewPasswordEntry()
	contenedorPass := container.NewVBox(layout.NewSpacer(), labelPass, TextPass, layout.NewSpacer())

	Btnlogin := widget.NewButton("Login", func() {

		if DB.RegistroExiste(TextUser.Text, DB.ExisteUsuario) {
			if Controles.VerificaPasword(DB.Buscarhass(TextUser.Text), TextPass.Text) {
				if DB.RolUser(TextUser.Text) == "Administrador" {
					VentanaAdmin(a, "Administrador", TextUser.Text)
					w.Hide()
				}
				if DB.RolUser(TextUser.Text) == "General" {
					VentanaUsuarios(a, DB.RolUser(TextUser.Text), TextUser.Text)
					Controles.RevisarUser(TextUser.Text)
					w.Hide()
				}
				TextUser.SetText("")
				TextPass.SetText("")
			} else {
				ShowError("error de contraseña", w)
			}
		} else {
			ShowError("error de usuario", w)
		}

	})

	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeyReturn {
			Btnlogin.OnTapped()
		}
	})
	contenedorBtnLogin := container.NewHBox(layout.NewSpacer(), Btnlogin, layout.NewSpacer())

	contenido := container.NewVBox(contenedorClock, contenedorTitulo, ContenedorUser, contenedorPass, contenedorBtnLogin)

	// Agregar los layouts a la ventana
	w.SetContent(contenido)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.Show()
	w.SetOnClosed(func() { a.Quit() })
}
