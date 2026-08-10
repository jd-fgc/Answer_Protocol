package main

import (
	"encoding/json"
	// "fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

var characterImages = map[string]string{
	"Isaac":  "client_gui/assets/Isaac.png",
	"Eve":    "client_gui/assets/Eve.png",
	"Judas":  "client_gui/assets/Judas.png",
	"Guppy":  "client_gui/assets/Guppy.png",
	"Azazel": "client_gui/assets/Azazel.png",
	"Lost":   "client_gui/assets/Lost.png",
}

type whoPlayerInfo struct {
	Name      string `json:"name"`
	Character string `json:"character"`
}

func fetchRoomPlayers() ([]whoPlayerInfo, error) {
	lines, err := netClient.SendCommand("WHO")
	if err != nil || len(lines) == 0 {
		return nil, err
	}
	var who struct {
		Room   []whoPlayerInfo `json:"room"`
		Server int             `json:"server"`
	}
	if err := json.Unmarshal([]byte(lines[0]), &who); err != nil {
		return nil, err
	}
	return who.Room, nil
}

func buildPlayerCard(info playerInfo) fyne.CanvasObject {
	var n_player Player
	n_player.Name = info.Name
	n_player.CURRENT_HP = info.HP
	n_player.Character = info.Character
	n_player.ATK = info.ATK
	n_player.COIN = info.COIN
	n_player.BOMB = info.BOMB
	n_player.KEY = info.KEY
	pv := NewPlayerView(&n_player)
	return pv.Container
}

func refreshRoomPlayers(leftContainer *fyne.Container, infos []playerInfo) {
	if len(infos) > 4 {
		infos = infos[:4]
	}

	leftContainer.RemoveAll()
	gap := float32(15)
	for i, info := range infos {
		leftContainer.Add(buildPlayerCard(info))
		if i < len(infos)-1 {
			spacer := canvas.NewRectangle(color.Transparent)
			spacer.SetMinSize(fyne.NewSize(1, gap))
			leftContainer.Add(spacer)
		}
	}
	leftContainer.Refresh()
}
