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
				res := MOVE(&usr, rooms, user_choice, &online_player)
				fmt.Println(res)
			case "LOOK":
				res, err := LOOK(usr, rooms, online_player)
				if err != nil {
					fmt.Println(err)
				}
				if res != "" {
					fmt.Println(res)
				}
			case "TALK":
				scanner.Scan()
				user_choice = scanner.Text()
				res := TALK(usr.Zone, rooms, user_choice)
				fmt.Println(res)
			case "TAKE":
				scanner.Scan()
				user_choice = scanner.Text()
				res := TAKE(&usr, &rooms, &online_player, user_choice)
				fmt.Println(res)
			case "DROP":
				scanner.Scan()
				user_choice = scanner.Text()
				res := DROP(&usr, &rooms, &online_player, user_choice)
				fmt.Println(res)
			case "\n":
				return nil
		}
	}
}
