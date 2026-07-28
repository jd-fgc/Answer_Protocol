package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_user_input(usr Player, rooms Rooms, online_player *Players, user_input string) error {
	scanner := bufio.NewScanner(os.Stdin)
	var user_choice string
	switch user_input {
	case "CONNECT":
		fmt.Print("DEBUG: Connect")
		scanner.Scan()
		user_choice = scanner.Text()
		choose_your_character(playable_characters)
		scanner.Scan()
		tmp_usr, err := character_selection(playable_characters, scanner.Text())
		if err != nil {
			return err
		}
		tmp_usr.Name = user_choice
		online_player.Players = append(online_player.Players, tmp_usr)
	case "MOVE":
		scanner.Scan()
		user_choice = scanner.Text()
		err := MOVE(&usr, rooms, user_choice, online_player)
		if err != nil {
			return err
		}
	case "LOOK":
		res, err := LOOK(usr, rooms, *online_player)
		if err != nil {
			fmt.Println(err)
			return err
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
		return nil
	}
	return nil
}
