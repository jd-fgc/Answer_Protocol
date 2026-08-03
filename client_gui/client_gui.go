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
var rooms Rooms
var visitedRooms = make(map[string]bool)

// FONCTION INUTILE
func load_quest() {
	quests.Quest = append(quests.Quest, Quest{Name: "Boss", Description: "Defeat the boss"})
}

func set_npc() {
	npc = append(npc, NPC{Name: "JD", Talk: "Salam les rhoyas", Zone: "test"})
	npc = append(npc, NPC{Name: "Toujours JD", Talk: "Salam les rhoyas", Zone: "test", Quest: "Boss"})
}

var playerViews []*PlayerView
var minimapContainer *fyne.Container

func fill_rooms() {
	rooms.Rooms = []Room{
		{Name: "test", East: "east1", South: "south1", Items: []string{"Potion", "Sword"}},
		{Name: "east1", West: "test", East: "east2"},
		{Name: "east2", West: "east1", Items: []string{"Key"}},
		{Name: "south1", North: "test", East: "south1east"},
		{Name: "south1east", West: "south1"},
	}
}

func refreshMiniMap() {
	minimapContainer.RemoveAll()
	minimapContainer.Add(buildMiniMap(rooms, player.zone, visitedRooms))
	minimapContainer.Refresh()
}

func fill_players() {
	players.Player = append(players.Player, PLAYER{NAME: "Isaac", Image: "client_gui/assets/Isaac.png", HP: 100, ATK: 3.5, COIN: 0, BOMB: 0, zone: "test"})
	players.Player = append(players.Player, PLAYER{NAME: "Eve", Image: "client_gui/assets/Eve.png", HP: 100, ATK: 3.5, COIN: 0, BOMB: 0, zone: "test"})
	players.Player = append(players.Player, PLAYER{NAME: "Judas", Image: "client_gui/assets/Judas.png", HP: 100, ATK: 3.5, COIN: 0, BOMB: 0, zone: "test"})
	players.Player = append(players.Player, PLAYER{NAME: "Guppy", Image: "client_gui/assets/Guppy.png", HP: 100, ATK: 3.5, COIN: 0, BOMB: 0, zone: "test"})
}

// CREATION D'UN OBJET PLAYER POUR CONTAINER
func NewPlayerView(p *PLAYER) *PlayerView {
	portrait := canvas.NewImageFromFile(p.Image)
	portrait.FillMode = canvas.ImageFillContain
	portrait.ScaleMode = canvas.ImageScalePixels
	portrait.SetMinSize(fyne.NewSize(247, 208))

	pv := &PlayerView{
		Player:   p,
		Portrait: portrait,
		Name:     widget.NewLabel(""),
		HP:       widget.NewLabel(""),
		ATK:      widget.NewLabel(""),
		Coin:     widget.NewLabel(""),
		Bomb:     widget.NewLabel(""),
	}

	infoColumn := container.NewVBox(
		pv.Name,
		pv.HP,
		pv.ATK,
		pv.Coin,
		pv.Bomb,
	)

	row := container.NewHBox(pv.Portrait, infoColumn)

	pv.Container = container.NewCenter(row)

	pv.Refresh()

	return pv
}

// FONCTION RESERVEE AUX OBJETS PlayerView
func (pv *PlayerView) Refresh() {
	p := pv.Player

	pv.Portrait.File = p.Image
	pv.Portrait.Refresh()

	pv.Name.SetText(p.NAME)
	pv.HP.SetText(fmt.Sprintf("HP : %d", p.HP))
	pv.ATK.SetText(fmt.Sprintf("ATK : %.1f", p.ATK))
	pv.Coin.SetText(fmt.Sprintf("Coins : %d", p.COIN))
	pv.Bomb.SetText(fmt.Sprintf("Bombs : %d", p.BOMB))
}

// LE MAIN
func main() {
	a := app.NewWithID("gleubeuskeufeu")
	a.Settings().SetTheme(&tapTheme{})
	w := a.NewWindow("The TCP of Isaac")
	w.Resize(fyne.NewSize(1600, 900))

	home := buildHomeScreen(func(username, ip, port string) {
		// TODO: à terme, la vraie réponse CONNECT du serveur
		charSelect := buildCharacterSelectScreen(username, func(character string) {
			startGame(w, username, character)
		})
		w.SetContent(charSelect)
	})

	w.SetContent(home)
	w.Show()
	a.Run()
}

func startGame(w fyne.Window, username string, character string) {
	actions = container.NewVBox()
	inventory = container.NewVBox()
	combat = container.NewVBox()
	log = container.NewVBox()
	log_text = widget.NewLabel("")
	fill_players()
	set_npc()
	load_quest()
	fill_rooms()

	player = players.Player[0]
	player.NAME = username
	player.zone = "test"
	visitedRooms[player.zone] = true

	views := []*PlayerView{
		NewPlayerView(&players.Player[0]),
		NewPlayerView(&players.Player[1]),
		NewPlayerView(&players.Player[2]),
		NewPlayerView(&players.Player[3]),
	}
	playerViews = views

	leftColumn := set_left_container(views)
	rightColumn := set_right_container(views)

	minimapContainer = container.NewStack()
	refreshMiniMap()

	bg := canvas.NewImageFromFile("client_gui/assets/backgroundv2.png")
	bg.FillMode = canvas.ImageFillStretch

	// Ajustement des blocs
	leftBlock := fixedSize(fyne.NewSize(437, 870), leftColumn)
	centerBlock := fixedSize(fyne.NewSize(557, 900), minimapContainer)
	rightBlock := fixedSize(fyne.NewSize(580, 870), rightColumn)

	leftBlock.Resize(fyne.NewSize(437, 870))
	leftBlock.Move(fyne.NewPos(0, 0))

	centerBlock.Resize(fyne.NewSize(557, 900))
	centerBlock.Move(fyne.NewPos(437, 0))

	rightOffset := float32(970)
	rightBlock.Resize(fyne.NewSize(580, 870))
	rightBlock.Move(fyne.NewPos(rightOffset, 0))

	gameLayout := container.NewWithoutLayout(leftBlock, centerBlock, rightBlock)

	content := container.NewStack(bg, gameLayout)
	w.SetContent(content)
}
