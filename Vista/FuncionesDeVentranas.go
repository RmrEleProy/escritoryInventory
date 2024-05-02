package Vista

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Genera una ventana de fyne recibe los parametros
// fyne.App, titulo,ancho y alto, y si se fija o no el tamaño de la ventana
func GenerarVentana(a fyne.App, titulo string, height, weidth float32, fixedsize bool) fyne.Window {
	w := a.NewWindow(titulo)
	w.Resize(fyne.NewSize(height, weidth))
	w.CenterOnScreen()
	w.SetFixedSize(fixedsize)
	return w
}

// muestra una ventana popup de algun error
func ShowError(message string, w fyne.Window) {
	dialog := dialog.NewError(fmt.Errorf(message), w)
	dialog.Show()
}

// ventana para validacion de la informacion
func MensajeValidacin(w fyne.Window, titulo, msg string) (respuesta string) {

	ventana := dialog.NewConfirm(titulo, msg, func(b bool) {
		if b {
			respuesta = "ok"
		} else {
			respuesta = "cancel"
		}
	}, w)
	ventana.Show()
	return respuesta
}

// ventana que informa al usuario
func MensajeInformacion(w fyne.Window, titulo, msg string) {
	ventana := dialog.NewInformation(titulo, msg, w)
	ventana.Show()
}

// Actualiza el relog
func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

// carga el logo de una ventana
func loadLogo(direccion string) fyne.Resource {
	data, err := os.ReadFile(direccion)
	if err != nil {
		panic(err)
	}
	logo := fyne.NewStaticResource("Logo", data)
	return logo
}

// resueleve la direccion para cargar los link u enlaces
func GetUrl(dir string) *url.URL {
	direccion, err := url.Parse(dir)
	if err != nil {
		panic(err)
	}
	return direccion
}
