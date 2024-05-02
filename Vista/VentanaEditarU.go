package Vista

import (
	"Inventario/DB"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func VentanaEditarU(a fyne.App, user DB.Usuarios) {
	w := GenerarVentana(a, "Editar usuario", 500, 450, true)
	labelNombre := widget.NewLabel("Nombre")
	entryNombre := widget.NewEntry()
	entryNombre.SetText(user.Nombre)

	labelEmail := widget.NewLabel("Email")
	entryEmail := widget.NewEntry()
	entryEmail.SetText(user.Email)

	labelContraseña := widget.NewLabel("Contraseña")
	entryContraseña := widget.NewEntry()

	labelRol := widget.NewLabel("Rol")
	Options := []string{"General", "Administrador"}
	selecRol := widget.NewSelect(Options, func(seleccionado string) {
	})

	BtnActualizar := widget.NewButton("Actualizar", func() {
		user.Nombre = entryNombre.Text
		user.Email = entryEmail.Text
		user.Contraseña = entryContraseña.Text
		user.Rol = selecRol.Selected
		user.Fecha_modificacion = time.Now().Format("02/01/2006")
		dialog.ShowConfirm("Confirmar", "¿Seguro que desea modificar este usuario?", func(b bool) {
			if b {
				DB.EditarUsuario(user)
				ListarUsuarios(a)
				w.Close()
			}
		}, w)
		// VentanaInventarios(a, "Administrador")
		// w.Close()
	})

	contenido := container.NewVBox(labelNombre, entryNombre, labelEmail, entryEmail, labelContraseña, entryContraseña, labelRol, selecRol, BtnActualizar)

	w.SetContent(contenido)
	w.Show()
}
