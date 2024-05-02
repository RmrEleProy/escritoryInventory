package Vista

import (
	"Inventario/DB"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func VentanaInventarios(a fyne.App, rol string) {
	w := GenerarVentana(a, "INVENTARIO", 600, 650, true)

	datos := DB.VerTodosLosRegistros()
	// Crear una tabla personalizada con encabezados
	inventario := widget.NewTable(
		func() (int, int) { return len(datos) + 1, 5 }, // Longitud de los datos + 1 para el encabezado
		func() fyne.CanvasObject { return widget.NewLabel("Placeholder") },
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			if id.Row == 0 {
				// Encabezados
				label := cell.(*widget.Label)
				switch id.Col {
				case 0:
					label.SetText("CODIGO")
				case 1:
					label.SetText("PRODUCTO")
				case 2:
					label.SetText("PRECIO")
				case 3:
					label.SetText("STOCK")
				case 4:
					label.SetText("PROVEEDOR")
				}
			} else if id.Row <= len(datos) {
				// Datos reales
				pro := datos[id.Row-1]
				label := cell.(*widget.Label)
				switch id.Col {
				case 0:
					label.SetText(fmt.Sprintf("%d", pro.Codigo))
				case 1:
					label.SetText(pro.Nombre)
				case 2:
					label.SetText(fmt.Sprintf("%.2f", pro.Precio))
				case 3:
					label.SetText(fmt.Sprintf("%d", pro.Stock))
				case 4:
					label.SetText(pro.Proveedor)
				}
			}
		},
	)

	inventario.OnSelected = func(id widget.TableCellID) {
		if id.Row > 0 {
			pro := datos[id.Row-1]
			if id.Col == 0 {
				if rol == "Administrador" {
					menuItems := []*fyne.MenuItem{
						fyne.NewMenuItem("Editar", func() { EditarProductosVentana(a, pro) }), //crear ventana para editar
						fyne.NewMenuItem("Eliminar", func() { DB.BorrarOldRegistro(pro.Id) }),
					}
					popupMenu := widget.NewPopUpMenu(fyne.NewMenu("", menuItems...), w.Canvas())
					popupMenu.ShowAtPosition(fyne.CurrentApp().Driver().AbsolutePositionForObject(inventario))
				}
			}
		}
	}

	w.SetContent(inventario)
	w.Show()
}
