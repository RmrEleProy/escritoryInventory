package Vista

import (
	"Inventario/Controles"
	"Inventario/DB"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// agrega un nuevo Usuario a la base de datos
func NewUser(a fyne.App) {
	w := GenerarVentana(a, "Nuevo Usuario", 400, 600, true)

	Lnombre := widget.NewLabel("Nombre")
	Enombre := widget.NewEntry()
	Lcorreo := widget.NewLabel("Correo")
	Ecorreo := widget.NewEntry()
	Lcontraseña := widget.NewLabel("Contraseña")
	Econtraseña := widget.NewPasswordEntry()
	Lrol := widget.NewLabel("Rol")
	Options := []string{"General", "Administrador"}
	Srol := widget.NewSelect(Options, func(seleccionado string) {
	})
	opciones := []string{"Fecha del sistema", "Fecha manual"}
	radio := widget.NewRadioGroup(opciones, nil)
	Lfecha := widget.NewLabel("Fecha inicial")
	Efecha := widget.NewEntry()
	Efecha.SetPlaceHolder("DD/MM/YYYY")
	radio.OnChanged = func(s string) {
		if s == "Fecha del sistema" {
			Efecha.SetText(time.Now().Format("02/01/2006"))
			Efecha.Disable()
		} else {
			Efecha.SetText("")
			Efecha.Enable()
		}
	}

	Lmsg := widget.NewLabel("")

	btnGuardar := widget.NewButton("Guardar", func() {
		user := DB.Usuarios{
			Nombre:             Enombre.Text,
			Email:              Ecorreo.Text,
			Contraseña:         Econtraseña.Text,
			Rol:                Srol.Selected,
			Fecha_creacion:     Efecha.Text,
			Fecha_modificacion: Efecha.Text,
			EstadoSesion:       "No Activo"}

		if Controles.ValidarCamposNewUser(user) != "ok" {
			ShowError(Controles.ValidarCamposNewUser(user), w)
		} else {
			nuevoUsuario := DB.Usuarios{
				Nombre:             Enombre.Text,
				Email:              Ecorreo.Text,
				Contraseña:         Controles.Encriptar(Econtraseña.Text),
				Rol:                Srol.Selected,
				Fecha_creacion:     Efecha.Text,
				Fecha_modificacion: Efecha.Text,
				EstadoSesion:       "No Activo"}

			dialog.ShowConfirm("Confirmar", "¿Seguro que desea agregar este usuario?", func(b bool) {
				if b {
					time.AfterFunc(1*time.Second, func() {
						Lmsg.SetText(DB.InsertarNewUser(nuevoUsuario))
						time.AfterFunc(3*time.Second, func() {
							VerTablaDeUsurios(a)
							Lmsg.SetText("")
						})
					})
					Enombre.SetText("")
					Ecorreo.SetText("")
					Econtraseña.SetText("")
					Srol.ClearSelected()
					Efecha.SetText("")
				}
			}, w)
		}
	})

	contenido := container.NewVBox(
		Lnombre,
		Enombre,
		Lcorreo,
		Ecorreo,
		Lcontraseña,
		Econtraseña,
		Lrol,
		Srol,
		radio,
		Lfecha,
		Efecha,
		btnGuardar,
		Lmsg)
	w.SetContent(contenido)
	w.Show()
}
