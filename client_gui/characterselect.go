package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

type characterOption struct {
	Name  string
	Image string
}

var characterOptions = []characterOption{
	{Name: "Isaac", Image: "client_gui/assets/Isaac.png"},
	{Name: "Eve", Image: "client_gui/assets/Eve.png"},
	{Name: "Judas", Image: "client_gui/assets/Judas.png"},
	{Name: "Guppy", Image: "client_gui/assets/Guppy.png"},
	{Name: "Azazel", Image: "client_gui/assets/Azazel.png"},
	{Name: "Lost", Image: "client_gui/assets/Lost.png"},
}

func buildCharacterSelectScreen(username string, onChoose func(character string)) fyne.CanvasObject {
	title := canvas.NewText("Choisis ton personnage, "+username, theme.Color(theme.ColorNameForeground))
	title.TextSize = 32
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	var cards []fyne.CanvasObject
	for _, c := range characterOptions {
		c := c
		tapImg := newTappableImage(c.Image, fyne.NewSize(220, 220), func() {
			onChoose(c.Name)
		})

		cards = append(cards, tapImg)
	}

	grid := container.NewGridWithColumns(3, cards...)

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		grid,
		layout.NewSpacer(),
	)

	return container.NewCenter(content)
}
