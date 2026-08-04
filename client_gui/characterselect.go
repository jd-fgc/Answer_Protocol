package main

import (
	"encoding/json"

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

func buildCharacterSelectScreen(username string, characterListJSON string, onChoose func(character string)) fyne.CanvasObject {
	var names []string
	if err := json.Unmarshal([]byte(characterListJSON), &names); err != nil {
		names = []string{"Isaac"}
	}

	title := canvas.NewText("Choisis ton personnage, "+username, theme.Color(theme.ColorNameForeground))
	title.TextSize = 32
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	imageFor := func(name string) string {
		return "client_gui/assets/" + name + ".png"
	}

	var cards []fyne.CanvasObject
	for _, name := range names {
		name := name
		tapImg := newTappableImage(imageFor(name), fyne.NewSize(220, 220), func() {
			onChoose(name)
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
