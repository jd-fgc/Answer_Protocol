package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type NPC struct {
	Name  string `json:"name"`
	Talk  string `json:"talk"`
	Zone  string `json:"room"`
	Quest string `json:"quest"`
}

type Player struct {
	Name      string
	Character string // nom du template choisi (Isaac, Eve...) — sert à retrouver le portrait
	HP        int64
	ATK       float64
	COIN      int64
	BOMB      int64
	Inventory []string
	Zone      string
}

type Players struct {
	Player []Player
}

type PlayerView struct {
	Player   *Player
	Portrait *canvas.Image
	Name     *widget.Label
	HP       *widget.Label
	ATK      *widget.Label
	Coin     *widget.Label
	Bomb     *widget.Label

	Container *fyne.Container
}

type Quests struct {
	Quest []Quest `json:"quests"`
}

type Quest struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	taken       bool
}

type ANSWER struct {
	cmd string
	res string
	err error
}

type Room struct {
	Name  string   `json:"name"`
	NPC   []NPC    `json:"NPC"`
	Items []string `json:"item"`
	North string   `json:"North"`
	South string   `json:"South"`
	East  string   `json:"East"`
	West  string   `json:"West"`
}

type Rooms struct {
	Rooms []Room `json:"world"`
}
