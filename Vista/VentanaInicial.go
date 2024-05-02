package Vista

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Inicio() {
	a := app.New()
	w := GenerarVentana(a, " ", 400, 450, true)

	clock := widget.NewLabel("")
	updateTime(clock)
	contenedorClock := container.NewHBox(layout.NewSpacer(), clock)

	menu1 := fyne.NewMenu("Menu",
		fyne.NewMenuItem("Login", func() { Login(a) }),
		fyne.NewMenuItem("acerca de", func() {}),
	)

	menu2 := fyne.NewMenu("Submenu 2",
		fyne.NewMenuItem("ayuda", func() {}),
		fyne.NewMenuItem("version", func() {}),
	)

	mainMenu := fyne.NewMainMenu(menu1, menu2)

	w.SetMainMenu(mainMenu)

	icono := loadLogo("static/IMG/icono.png")
	w.SetIcon(icono)

	fotoPerfil := canvas.NewImageFromFile("static/IMG/Designer.png")
	fotoPerfil.FillMode = canvas.ImageFillOriginal

	iconohub := widget.NewIcon(loadLogo("static/IMG/Iconos/github.png"))
	// linkGithub, _ := url.Parse("https://github.com/RmrEleProy")
	linkhub := widget.NewHyperlink("Mis proyectos", GetUrl("https://github.com/RmrEleProy"))

	correo := widget.NewIcon(theme.MailComposeIcon())

	contenedorlinks := container.NewHBox(iconohub, linkhub, layout.NewSpacer(), correo)

	contenido := container.NewVBox(contenedorClock, fotoPerfil, contenedorlinks)
	w.SetContent(contenido)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()

}
