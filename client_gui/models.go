package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type NPC struct {
	Name   string `json:"name"`
	Talk   string `json:"talk"`
	HP     int    `json:"hp"`
	Ennemy bool   `json:"ennemy"`
	Zone   string `json:"room"`
	ATK    int    `json:"attack"`
	Quest  string `json:"quest"`
}

type Player struct {
	Name             string
	Character        string
	Item             string
	Inventory        []string
	MAX_HP           int
	CURRENT_HP       int
	ATK              float32
	DMG_Mult         float32
	Luck             int
	Deal_chance      int
	Zone             string
	Damage_Reduction float32
	Quests           []Quest
	COIN             int
	BOMB             int
	KEY              int
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
	Name          string `json:"name"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	Completed     bool   `json:"completed"`
	Reward        string `json:"reward"`
	Amount_reward int    `json:"amount_reward"`
	Item_needed   string `json:"item_needed"`
	Needed_amount int    `json:"needed_amount"`
	Ennemy        string `json:"ennemy"`
	taken         bool
	finished      bool
}

type Room struct {
	Name  string   `json:"name"`
	NPC   []NPC    `json:"NPC"`
	Items []string `json:"item"`
	North string   `json:"North"`
	South string   `json:"South"`
	East  string   `json:"East"`
	West  string   `json:"West"`
	Type  string   `json:"Type"`
}

type Rooms struct {
	Rooms []Room `json:"world"`
}
