package Vista

import (
	"Inventario/DB"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Muestra los usuarios en la base de datos
func VerTablaDeUsurios(a fyne.App) {
	w := GenerarVentana(a, "Personas en la base de datos", 1024, 650, false)
	datos := DB.VerTodosLosUsuarios()

	lista := widget.NewList(
		func() int {
			return len(datos)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("NOMBRE"),
				widget.NewLabel("CORREO"),
				widget.NewLabel("ROL"),
				widget.NewLabel("FECHA DE CREACION"),
				widget.NewLabel("FECHA DE ULTIMA MODIFICACION"),
			)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			t := item.(*fyne.Container)
			user := datos[id]
			t.Objects[0].(*widget.Label).SetText(user.Nombre)
			t.Objects[1].(*widget.Label).SetText(user.Email)
			t.Objects[2].(*widget.Label).SetText(user.Rol)
			t.Objects[3].(*widget.Label).SetText(user.Fecha_creacion)
			t.Objects[4].(*widget.Label).SetText(user.Fecha_modificacion)
		},
	)
	lista.OnSelected = func(id widget.ListItemID) {
		user := datos[id]
		menuItems := []*fyne.MenuItem{
			fyne.NewMenuItem("Editar", func() { VentanaEditarU(a, user) }),
			fyne.NewMenuItem("Eliminar", func() { DB.DeleteUserbyId(user.Id) })}
		popupMenu := widget.NewPopUpMenu(fyne.NewMenu("", menuItems...), w.Canvas())
		popupMenu.ShowAtPosition(fyne.CurrentApp().Driver().AbsolutePositionForObject(lista))
	}

	w.SetContent(lista)
	w.Show()
}

func ListarUsuarios(a fyne.App) {
	w := GenerarVentana(a, "Personas en la base de datos", 1024, 650, false)
	lista := actualizarLista(w, a)
	w.SetContent(lista)
	w.Show()
}

func actualizarLista(w fyne.Window, a fyne.App) *widget.List {
	datos := DB.VerTodosLosUsuarios()
	lista := widget.NewList(
		func() int {
			return len(datos)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("NOMBRE"),
				widget.NewLabel("CORREO"),
				widget.NewLabel("ROL"),
				widget.NewLabel("FECHA DE CREACION"),
				widget.NewLabel("FECHA DE ULTIMA MODIFICACION"),
			)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			t := item.(*fyne.Container)
			user := datos[id]
			t.Objects[0].(*widget.Label).SetText(user.Nombre)
			t.Objects[1].(*widget.Label).SetText(user.Email)
			t.Objects[2].(*widget.Label).SetText(user.Rol)
			t.Objects[3].(*widget.Label).SetText(user.Fecha_creacion)
			t.Objects[4].(*widget.Label).SetText(user.Fecha_modificacion)
		},
	)
	lista.OnSelected = func(id widget.ListItemID) {
		user := datos[id]
		menuItems := []*fyne.MenuItem{
			fyne.NewMenuItem("Editar", func() { VentanaEditarU(a, user) }),
			fyne.NewMenuItem("Eliminar", func() {
				dialog.ShowConfirm("Confirmar", "¿Seguro que desea eliminar este usuario?", func(b bool) {
					if b {
						DB.DeleteUserbyId(user.Id)
						lista = actualizarLista(w, a)
						w.SetContent(lista)
					}
				}, w)
			})}
		popupMenu := widget.NewPopUpMenu(fyne.NewMenu("", menuItems...), w.Canvas())
		popupMenu.ShowAtPosition(fyne.CurrentApp().Driver().AbsolutePositionForObject(lista))
	}
	return lista
}
