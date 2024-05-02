package Vista

import (
	"Inventario/DB"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ventana de edicion para el producto seleccionados
func EditarProductosVentana(a fyne.App, p DB.ProStru) {
	w := GenerarVentana(a, "EDITAR PRODUCTOS", 400, 450, true)

	labelCodigo := widget.NewLabel("Codigo")
	entryCodigo := widget.NewEntry()
	entryCodigo.SetText(strconv.Itoa(p.Codigo))

	labelProducto := widget.NewLabel("Producto")
	entryProducto := widget.NewEntry()
	entryProducto.SetText(p.Nombre)

	lasbelcateggria := widget.NewLabel("Categoria")
	entrycategoria := widget.NewEntry()
	entrycategoria.SetText(p.Categoria)

	labelPrecio := widget.NewLabel("Precio")
	entryPrecio := widget.NewEntry()
	entryPrecio.SetText(strconv.FormatFloat(p.Precio, 'f', -1, 64))

	labelStock := widget.NewLabel("Stock")
	entryStock := widget.NewEntry()
	entryStock.SetText(strconv.Itoa(p.Stock))

	labelProveedor := widget.NewLabel("Proveedor")
	entryProveedor := widget.NewEntry()
	entryProveedor.SetText(p.Proveedor)

	BtnActualizar := widget.NewButton("Actualizar", func() {
		p.Codigo, _ = strconv.Atoi(entryCodigo.Text)
		p.Nombre = entryProducto.Text
		p.Categoria = entrycategoria.Text
		p.Precio, _ = strconv.ParseFloat(entryPrecio.Text, 64)
		p.Stock, _ = strconv.Atoi(entryStock.Text)
		p.Proveedor = entryProveedor.Text
		DB.EditarProducto(p)
		VentanaInventarios(a, "Administrador")
		w.Close()
	})

	contenido := container.NewVBox(labelCodigo, entryCodigo, labelProducto, entryProducto, lasbelcateggria, entrycategoria, labelPrecio, entryPrecio, labelStock, entryStock, labelProveedor, entryProveedor, BtnActualizar)

	w.SetContent(contenido)
	w.Show()
}
