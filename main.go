package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	// "io"
	"os"
	"bufio"
)


type Players struct {
	Players []Player
}


type Player struct {
	Name string
	Item string
	Inventory []string
	MAX_HP int
	CURRENT_HP int
	ATK float32
	DMG_Mult float32
	Luck int
	Deal_chance int
	Respawn int
	Zone string
}


type Rooms struct {
	Rooms []Room `json:"world"`
}


type Room struct {
	Name string `json:"name"`
	NPC []NPC `json:"NPC"`
	Items []string `json:"item"`
	North string `json:"North"`
	South string `json:"South"`
	East string `json:"East"`
	West string `json:"West"`
}


type NPC struct {
	Name string `json:"name"`
	Talk string `json:"talk"`
	HP int `json:"hp"`
	Ennemy bool `json:"ennemy"`
}


// CALL THIS FUNCTION EVERYTIME A PLAYER'S DATA IS MODIFIED
func update_players(player Player, players *Players) error {
	for i, val := range players.Players {
		if val.Name == player.Name {
			players.Players[i] = player
			return nil
		}
	}
	return errors.New("Player not found")
}


func main() {
	var user_input string
	var online_player Players
	scanner := bufio.NewScanner(os.Stdin)
	rooms, err := load_map("world_map.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	charac, err := read_character_file("Characters.json", rooms.Rooms[0].Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	choose_your_character(charac)
	fmt.Scan(&user_input)
	usr, err := character_selection(charac, user_input)
	online_player.Players = append(online_player.Players, usr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var user_choice string
		fmt.Scan(&user_input)
		switch user_input {
			case "MOVE":
				scanner.Scan()
				user_choice = scanner.Text()
				err := MOVE(&usr, rooms, user_choice, &online_player)
				if err != nil {
					fmt.Println(err)
					return
				}
			case "LOOK":
				res, err := LOOK(usr, rooms, online_player)
				if err != nil {
					fmt.Println(err)
					return
				}
				if res != "" {
					fmt.Println(res)
				}
			case "TALK":
				scanner.Scan()
				user_choice = scanner.Text()
				res := TALK(usr.Zone, rooms, user_choice)
				if res != "" {
					fmt.Println(res)
				}
			case "\n":
				return
		}
	}
}