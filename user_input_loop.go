package main

import (
	"fmt"
	"bufio"
	"os"
)


func get_user_input(usr Player) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var user_input string
		var user_choice string
		fmt.Scan(&user_input)
		switch user_input {
			case "MOVE":
				scanner.Scan()
				user_choice = scanner.Text()
				res := MOVE(&usr, user_choice)
				fmt.Println(res)
			case "LOOK":
				res, err := LOOK(usr)
				if err != nil {
					fmt.Println(err)
				}
				if res != "" {
					fmt.Println(res)
				}
			case "TALK":
				scanner.Scan()
				user_choice = scanner.Text()
				res := TALK(usr.Zone, user_choice)
				fmt.Println(res)
			case "TAKE":
				scanner.Scan()
				user_choice = scanner.Text()
				res := TAKE(&usr, user_choice)
				fmt.Println(res)
			case "DROP":
				scanner.Scan()
				user_choice = scanner.Text()
				res := DROP(&usr, user_choice)
				fmt.Println(res)
			case "INVENTORY":
				res, inv := INVENTORY(usr)
				fmt.Println(res, inv)
			case "ATTACK":
				scanner.Scan()
				user_choice = scanner.Text()
				res := ATTACK(usr, user_choice)
				fmt.Println(res)
			case "DEFEND":
				res := DEFEND(&usr)
				fmt.Println(res)
			case "STATUS":
				res := STATUS(usr)
				fmt.Println(res)
			case "QUEST":
				scanner.Scan()
				user_choice = scanner.Text()
				res := QUEST(&usr, user_choice)
				fmt.Println(res)
			case "QUESTS":
				res := QUESTS()
				fmt.Println(res)
			case "\n":
				return nil
		}
	}
}
