package main

type Players struct {
	Players []Player
}

type Player struct {
	Name        string
	Character   string
	Item        string
	Inventory   []string
	MAX_HP      int
	CURRENT_HP  int
	ATK         float32
	DMG_Mult    float32
	Luck        int
	Deal_chance int
	Respawn     int
	Zone        string
}

type Rooms struct {
	Rooms []Room `json:"world"`
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

type NPC struct {
	Name   string `json:"name"`
	Talk   string `json:"talk"`
	HP     int    `json:"hp"`
	Ennemy bool   `json:"ennemy"`
}
