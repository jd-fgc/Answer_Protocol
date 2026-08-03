package main

import (
	"encoding/json"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func fetchRoomPlayers() ([]string, error) {
	lines, err := netClient.SendCommand("WHO")
	if err != nil || len(lines) == 0 {
		return nil, err
	}

	var who struct {
		Room   []string `json:"room"`
		Server int      `json:"server"`
	}
	if err := json.Unmarshal([]byte(lines[0]), &who); err != nil {
		return nil, err
	}
	return who.Room, nil
}

func buildPlayerCard(name string) fyne.CanvasObject {
	if name == player.NAME {
		pv := NewPlayerView(&player)
		return pv.Container
	}

	placeholder := canvas.NewRectangle(color.NRGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff})
	placeholder.SetMinSize(fyne.NewSize(247, 208))

	nameLabel := widget.NewLabel(name)
	row := container.NewHBox(placeholder, nameLabel)
	return container.NewCenter(row)
}

func refreshRoomPlayers(leftContainer *fyne.Container) {
	names, err := fetchRoomPlayers()
	if err != nil {
		return
	}
	if len(names) > 4 {
		names = names[:4]
	}

	leftContainer.RemoveAll()
	gap := float32(26)

	for i, name := range names {
		leftContainer.Add(buildPlayerCard(name))
		if i < len(names)-1 {
			spacer := canvas.NewRectangle(color.Transparent)
			spacer.SetMinSize(fyne.NewSize(1, gap))
			leftContainer.Add(spacer)
		}
	}
	leftContainer.Refresh()
}
