package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func buildHomeScreen(onConnect func(username string, characterListJSON string)) fyne.CanvasObject {
	bg := canvas.NewImageFromFile("client_gui/assets/title_screen.png")
	bg.FillMode = canvas.ImageFillStretch

	widthSpacer := canvas.NewRectangle(color.Transparent)
	widthSpacer.SetMinSize(fyne.NewSize(420, 1))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Pseudo")

	ipEntry := widget.NewEntry()
	ipEntry.SetText("localhost")

	portEntry := widget.NewEntry()
	portEntry.SetText("9090")

	status := widget.NewLabel("")

	connectBtn := widget.NewButton("CONNECT", func() {
		if nameEntry.Text == "" {
			status.SetText("Choisis un pseudo avant de te connecter.")
			return
		}
		status.SetText("Connexion en cours...")

		lines, err := doConnect(ipEntry.Text, portEntry.Text, nameEntry.Text)
		if err != nil {
			status.SetText("Erreur : " + err.Error())
			return
		}

		if len(lines) < 2 {
			status.SetText("Réponse serveur incomplète")
			return
		}

		onConnect(nameEntry.Text, lines[1])
	})
	connectBtn.Importance = widget.HighImportance

	form := container.NewVBox(
		widthSpacer,
		nameEntry,
		ipEntry,
		portEntry,
		connectBtn,
		status,
	)

	topPush := canvas.NewRectangle(color.Transparent)
	topPush.SetMinSize(fyne.NewSize(1, 800))

	overlay := container.NewVBox(
		topPush,
		container.NewHBox(layout.NewSpacer(), form, layout.NewSpacer()),
		layout.NewSpacer(),
	)

	return container.NewStack(bg, overlay)
}
