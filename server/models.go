package main

type Players struct {
	Players []Player
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
	Respawn          int
	Zone             string
	Damage_Reduction float32
	Quests           []Quest
	Coin             int
	Bomb             int
	Key              int
}

type Rooms struct {
	Rooms []Room `json:"world"`
}

type ItemInstance struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	Name  string         `json:"name"`
	NPC   []NPC          `json:"NPC"`
	Items []ItemInstance `json:"-"`
	North string         `json:"North"`
	South string         `json:"South"`
	East  string         `json:"East"`
	West  string         `json:"West"`
}

var itemDisplayName = make(map[string]string)

type NPC struct {
	Name   string `json:"name"`
	Talk   string `json:"talk"`
	HP     int    `json:"hp"`
	Ennemy bool   `json:"ennemy"`
	Zone   string `json:"room"`
	ATK    int    `json:"attack"`
	Quest  string `json:"quest"`
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
