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

func load_quests(file_path string) (Quests, error) {
	var q Quests
	file, err := os.Open(file_path)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &q)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	return q, nil
}

var playerViews []*PlayerView
var minimapContainer fyne.CanvasObject

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

var mapContentHolder *fyne.Container
var mapDraggableWidget *draggable

func refreshMiniMap() {
	gridObj, size := buildMiniMap(rooms, player.Zone, visitedRooms)
	mapContentHolder.RemoveAll()
	mapContentHolder.Add(gridObj)
	mapContentHolder.Refresh()

	mapDraggableWidget.Resize(size)
	mapDraggableWidget.Move(mapDraggableWidget.clamp(mapDraggableWidget.Position()))
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

var mapScrollGlobal *container.Scroll

func startGame(w fyne.Window, username string, character string) {
	actions = container.NewVBox()
	inventory = container.NewVBox()
	combat = container.NewVBox()
	log = container.NewVBox()
	log_text = widget.NewLabel("")
	log_text.Wrapping = fyne.TextWrapWord
	quests, err = load_quests("data/quests.json")
	if err != nil {
		return
	}

	var errChar error
	players, errChar = read_character_file("data/Characters.json", "")
	if errChar != nil {
		log_text.SetText("Erreur chargement personnages: " + errChar.Error())
	}

	player = findPlayerTemplate(character)
	player.Name = username

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

	// Tailles fixes des blocs
	leftW, leftH := float32(437), float32(870)
	centerW, centerH := float32(557), float32(900)
	rightW, rightH := float32(554), float32(963)

	mapContentHolder = container.NewStack()
	mapDraggableWidget = newDraggable(mapContentHolder, fyne.NewSize(centerW, centerH))
	minimapContainer = container.NewWithoutLayout(mapDraggableWidget)

	refreshRoomState(false)

	mapDraggableWidget.Move(fyne.NewPos(
		(centerW-mapDraggableWidget.Size().Width)/2,
		(centerH-mapDraggableWidget.Size().Height)/2,
	))

	bg := canvas.NewImageFromFile("client_gui/assets/backgroundv4.png")
	bg.FillMode = canvas.ImageFillStretch

	leftBlock := fixedSize(fyne.NewSize(leftW, leftH), leftColumn)
	centerBlock := fixedSize(fyne.NewSize(centerW, centerH), minimapContainer)
	rightBlock := fixedSize(fyne.NewSize(rightW, rightH), rightColumn)

	left_xoffset := float32(-10)
	left_yoffset := float32(73)
	leftBlock.Resize(fyne.NewSize(leftW, leftH))
	leftBlock.Move(fyne.NewPos(left_xoffset, left_yoffset))

	center_xoffset := left_xoffset + leftW
	center_yoffset := float32(0)
	centerBlock.Resize(fyne.NewSize(centerW, centerH))
	centerBlock.Move(fyne.NewPos(center_xoffset, center_yoffset))

	right_xOffset := float32(998)
	right_yOffset := float32(23)
	rightBlock.Resize(fyne.NewSize(rightW, rightH))
	rightBlock.Move(fyne.NewPos(right_xOffset, right_yOffset))

	gameLayout := container.NewWithoutLayout(leftBlock, centerBlock, rightBlock)

	content := container.NewStack(bg, gameLayout)
	w.SetContent(content)
}
