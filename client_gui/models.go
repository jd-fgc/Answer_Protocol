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

type PLAYER struct {
	NAME      string
	Image     string
	HP        int64
	ATK       float64
	COIN      int64
	BOMB      int64
	Inventory []string
	zone      string
}

type PLAYERS struct {
	Player []PLAYER
}

type PlayerView struct {
	Player *PLAYER

	Portrait *canvas.Image
	Name     *widget.Label
	HP       *widget.Label
	ATK      *widget.Label
	Coin     *widget.Label
	Bomb     *widget.Label

	Container *fyne.Container
}

type Quests struct {
	Quest []Quest
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
	Name  string
	NPC   []NPC
	Items []string
	North string
	South string
	East  string
	West  string
}

type Rooms struct {
	Rooms []Room
}
