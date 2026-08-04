package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var player Player
var players Players
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
var err error

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

func fill_rooms(filepath string) (Rooms, error) {
	var r Rooms
	file, err := os.Open(filepath)
	if err != nil {
		return Rooms{Rooms: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Rooms{Rooms: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &r)
	if err != nil {
		return Rooms{Rooms: nil}, errors.New(err.Error())
	}
	return r, nil
}

func refreshMiniMap() {
	minimapContainer.RemoveAll()
	minimapContainer.Add(buildMiniMap(rooms, player.Zone, visitedRooms))
	minimapContainer.Refresh()
}

func read_character_file(filepath string, start_zone string) (Players, error) {
	var charac []Player
	file, err := os.Open(filepath)
	if err != nil {
		return Players{Player: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Players{Player: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &charac)
	if err != nil {
		return Players{Player: nil}, errors.New(err.Error())
	}
	for i, _ := range charac {
		charac[i].Zone = start_zone
		charac[i].Character = charac[i].Name
	}
	return Players{Player: charac}, nil
}

// CREATION D'UN OBJET PLAYER POUR CONTAINER
func NewPlayerView(p *Player) *PlayerView {
	portrait := canvas.NewImageFromFile(characterImages[p.Character])
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

	infoColumn := container.NewVBox(pv.Name, pv.HP, pv.ATK, pv.Coin, pv.Bomb)
	row := container.NewHBox(pv.Portrait, infoColumn)
	pv.Container = container.NewCenter(row)
	pv.Refresh()
	return pv
}

// FONCTION RESERVEE AUX OBJETS PlayerView
func (pv *PlayerView) Refresh() {
	p := pv.Player

	pv.Portrait.File = characterImages[p.Character]
	pv.Portrait.Refresh()

	pv.Name.SetText(p.Name)
	pv.HP.SetText(fmt.Sprintf("HP : %d", p.HP))
	pv.ATK.SetText(fmt.Sprintf("ATK : %.1f", p.ATK))
	pv.Coin.SetText(fmt.Sprintf("Coins : %d", p.COIN))
	pv.Bomb.SetText(fmt.Sprintf("Bombs : %d", p.BOMB))
}

func findPlayerTemplate(name string) Player {
	for _, p := range players.Player {
		if p.Name == name {
			return p
		}
	}
	return players.Player[0]
}

// LE MAIN
func main() {
	a := app.NewWithID("gleubeuskeufeu")
	a.Settings().SetTheme(&tapTheme{})
	w := a.NewWindow("The TCP of Isaac")
	w.Resize(fyne.NewSize(1600, 900))
	// w.SetFixedSize(true)

	home := buildHomeScreen(func(username string, characterListJSON string) {
		charSelect := buildCharacterSelectScreen(username, characterListJSON, func(character string) {
			lines, err := netClient.SendCommand("CHOOSE " + character)
			if err != nil || len(lines) == 0 {
				println("ERREUR : ", err)
				return
			}
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
	set_npc()
	load_quest()
	rooms, err = fill_rooms("data/world_map.json")
	if err != nil {
		return
	}

	players, err = read_character_file("data/Characters.json", rooms.Rooms[0].Name)
	player = findPlayerTemplate(character)
	player.Name = username
	visitedRooms[player.Zone] = true

	views := []*PlayerView{
		NewPlayerView(&players.Player[0]),
		NewPlayerView(&players.Player[1]),
		NewPlayerView(&players.Player[2]),
		NewPlayerView(&players.Player[3]),
	}
	playerViews = views
	chatTab := buildChatTab()
	groupTab := buildGroupTab()
	go listenChatEvents()

	leftColumn := set_left_container()
	leftColumnGlobal = leftColumn
	rightColumn := set_right_container(views, chatTab, groupTab)

	refreshRoomPlayers(leftColumn)

	minimapContainer = container.NewStack()
	refreshMiniMap()

	bg := canvas.NewImageFromFile("client_gui/assets/backgroundv2.png")
	bg.FillMode = canvas.ImageFillStretch

	// Ajustement des blocs
	leftBlock := fixedSize(fyne.NewSize(437, 870), leftColumn)
	centerBlock := fixedSize(fyne.NewSize(557, 900), minimapContainer)
	rightBlock := fixedSize(fyne.NewSize(555, 800), rightColumn)

	// bloc gauche
	left_xoffset := float32(-10)
	left_yoffset := float32(26)
	leftBlock.Resize(fyne.NewSize(437, 870))
	leftBlock.Move(fyne.NewPos(left_xoffset, left_yoffset))

	// bloc milieu
	centerBlock.Resize(fyne.NewSize(557, 900))
	centerBlock.Move(fyne.NewPos(280, -300))

	// bloc de droite
	right_xOffset := float32(998)
	right_yOffset := float32(23)

	rightBlock.Resize(fyne.NewSize(550, 965))
	rightBlock.Move(fyne.NewPos(right_xOffset, right_yOffset))

	gameLayout := container.NewWithoutLayout(leftBlock, centerBlock, rightBlock)

	content := container.NewStack(bg, gameLayout)
	w.SetContent(content)
}
