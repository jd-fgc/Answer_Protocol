package main

import (
	"fmt"
	"bufio"
	"os"
)


func get_user_input(usr Player, rooms Rooms, online_player Players) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var user_input string
		var user_choice string
		fmt.Scan(&user_input)
		switch user_input {
			case "MOVE":
				scanner.Scan()
				user_choice = scanner.Text()
				err := MOVE(&usr, rooms, user_choice, &online_player)
				if err != nil {
					return err
				}
			case "LOOK":
				res, err := LOOK(usr, rooms, online_player)
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
	}
}
