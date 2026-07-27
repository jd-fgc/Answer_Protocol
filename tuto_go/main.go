package main

import (
	"errors"
	"fmt"
)

type Room struct {
	R_Name string
	Items  []string
}

type Player struct {
	Name      string
	Inventory []string
}

func TakeItem(room *Room, player *Player, itemName string) error {
	for i, item := range room.Items {
		if item == itemName {
			room.Items = append(room.Items[:i], room.Items[i+1:]...)
			player.Inventory = append(player.Inventory, itemName)
			return nil
		}
	}
	return errors.New("Item not found: " + itemName)
}

func DropItem(room *Room, player *Player, itemName string) error {
	for i, item := range player.Inventory {
		if item == itemName {
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			room.Items = append(room.Items, itemName)
			return nil
		}
	}
	return errors.New("Item not found: " + itemName)
}

func main() {
	Entrance := Room{R_Name: "Entrance", Items: []string{"sword", "shield", "potion"}}
	jd := Player{Name: "JD", Inventory: []string{}}
	fmt.Print("Initialisation")
	fmt.Print("\nRoom name = ", Entrance.R_Name)
	fmt.Print("\nItems in Room", Entrance.Items)

	fmt.Print("\nPlayer name = ", jd.Name)
	fmt.Print("\nPlayer inventory = ", jd.Inventory)

	fmt.Print("\n")
	fmt.Print("\nJd prend une potion dans la room")
	err := TakeItem(&Entrance, &jd, "potion")
	if err != nil {
		fmt.Print("\nErreur: ", err)
	}
	fmt.Print("\nRoom name = ", Entrance.R_Name)
	fmt.Print("\nItems in Room", Entrance.Items)

	fmt.Print("\nPlayer name = ", jd.Name)
	fmt.Print("\nPlayer inventory = ", jd.Inventory)

	fmt.Print("\n")
	fmt.Print("\nJd lache une potion")
	err = DropItem(&Entrance, &jd, "potion")
	if err != nil {
		fmt.Print("\nErreur: ", err)
	}
	fmt.Print("\nRoom name = ", Entrance.R_Name)
	fmt.Print("\nItems in Room", Entrance.Items)

	fmt.Print("\nPlayer name = ", jd.Name)
	fmt.Print("\nPlayer inventory = ", jd.Inventory)

	fmt.Print("\n\nJd essaie de prendre une hache")
	err = TakeItem(&Entrance, &jd, "axe")
	if err != nil {
		fmt.Print("\nErreur: ", err)
	}
}
