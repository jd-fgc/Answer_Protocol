package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func update_players(player Player) error {
	for i, val := range online_players.Players {
		if val.Name == player.Name {
			online_players.Players[i] = player
			return nil
		}
	}
	return errors.New("Player not found")
}

func read_character_file(filepath string, start_zone string) (Players, error) {
	var charac []Player
	file, err := os.Open(filepath)
	if err != nil {
		return Players{Players: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Players{Players: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &charac)
	if err != nil {
		return Players{Players: nil}, errors.New(err.Error())
	}
	for i, _ := range charac {
		charac[i].Zone = start_zone
	}
	return Players{Players: charac}, nil
}

func load_map(filepath string) (Rooms, error) {
	var rooms Rooms
	file, err := os.Open(filepath)
	if err != nil {
		return Rooms{Rooms: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Rooms{Rooms: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &rooms)
	if err != nil {
		return Rooms{Rooms: nil}, errors.New(err.Error())
	}
	return rooms, nil
}

// func choose_your_character(characters Players) {
// 	fmt.Println("Choose one of the following character:")
// 	for _, val := range characters.Players {
// 		fmt.Println("----------")
// 		fmt.Println(val.Name)
// 	}
// 	fmt.Println("----------")
// }

func character_selection(characters Players, user_choice string) (Player, error) {
	for _, val := range characters.Players {
		if val.Name == user_choice {
			return val, nil
		}
	}
	return Player{}, errors.New("No player found with this name")
}
