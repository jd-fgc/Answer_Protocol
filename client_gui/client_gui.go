package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var player PLAYER
var players PLAYERS
var answer ANSWER
var npc []NPC
var actions *fyne.Container
var inventory *fyne.Container
var combat *fyne.Container
var log *fyne.Container
var log_text *widget.Label
var quests Quests

// FONCTION INUTILE
func load_quest() {
	quests.Quest = append(quests.Quest, Quest{Name: "Boss", Description: "Defeat the boss"})
}

func set_npc() {
	npc = append(npc, NPC{Name: "JD", Talk: "Salam les rhoyas", Zone: "test"})
	npc = append(npc, NPC{Name: "Toujours JD", Talk: "Salam les rhoyas", Zone: "test", Quest: "Boss"})
}

func fill_players() {
	players.Player = append(players.Player, PLAYER{NAME: "Boubou", HP: 100, ATK: 3.5, COIN: 0, BOMB: 0, zone: "test"})
	players.Player = append(players.Player, PLAYER{NAME: "Bebou", HP: 100, ATK: 3.5, COIN: 0, BOMB: 0, zone: "test"})
	players.Player = append(players.Player, PLAYER{NAME: "Babou", HP: 100, ATK: 3.5, COIN: 0, BOMB: 0, zone: "test"})
	players.Player = append(players.Player, PLAYER{NAME: "Bibou", HP: 100, ATK: 3.5, COIN: 0, BOMB: 0, zone: "test"})
}

// CREATION D'UN OBJET PLAYER POUR CONTAINER
func NewPlayerView(p *PLAYER) *PlayerView {
	pv := &PlayerView{
		Player: p,
		Name:   widget.NewLabel(""),
		HP:     widget.NewLabel(""),
		ATK:    widget.NewLabel(""),
		Coin:   widget.NewLabel(""),
		Bomb:   widget.NewLabel(""),
	}

	pv.Container = container.NewVBox(
		pv.Name,
		pv.HP,
		pv.ATK,
		pv.Coin,
		pv.Bomb,
	)

	pv.Refresh()

	return pv
}

// FONCTION RESERVEE AUX OBJETS PlayerView
func (pv *PlayerView) Refresh() {
	p := pv.Player

	pv.Name.SetText(p.NAME)
	pv.HP.SetText(fmt.Sprintf("HP : %d", p.HP))
	pv.ATK.SetText(fmt.Sprintf("ATK : %.1f", p.ATK))
	pv.Coin.SetText(fmt.Sprintf("Coins : %d", p.COIN))
	pv.Bomb.SetText(fmt.Sprintf("Bombs : %d", p.BOMB))
}

// LE MAIN
func main() {
	a := app.NewWithID("gleubeuskeufeu")
	w := a.NewWindow("The TCP of Isaac")
	w.Resize(fyne.NewSize(1600, 900))
	actions = container.NewVBox()
	inventory = container.NewVBox()
	combat = container.NewVBox()
	log = container.NewVBox()
	log_text = widget.NewLabel("")
	fill_players()
	set_npc()
	load_quest()
	player = players.Player[0]
	views := []*PlayerView{
		NewPlayerView(&players.Player[0]),
		NewPlayerView(&players.Player[1]),
		NewPlayerView(&players.Player[2]),
		NewPlayerView(&players.Player[3]),
	}
	img := canvas.NewImageFromFile("data/Title_screen.png")
	img.FillMode = canvas.ImageFillOriginal
	leftColumn := set_left_container(views)
	rightColumn := set_right_container(views)
	SCREEN := container.NewVBox(
		img,
	)
	content := container.NewHSplit(
		container.NewHSplit(leftColumn, SCREEN),
		rightColumn,
	)
	w.SetContent(content)
	w.Show()
	a.Run()
}
