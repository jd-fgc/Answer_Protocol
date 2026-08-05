package main

import (
	"strings"

	"fyne.io/fyne/v2/widget"
)

func look_room_description() string {
	r := current_room()
	var npcNames []string
	for _, n := range npc {
		if n.Zone == r.Name {
			npcNames = append(npcNames, n.Name)
		}
	}
	return "Room: " + r.Name +
		"\nItems: " + strings.Join(r.Items, ", ") +
		"\nNPCs: " + strings.Join(npcNames, ", ")
}

func inventory_description() string {
	if len(player.Inventory) == 0 {
		return "Inventory: (empty)"
	}
	return "Inventory: " + strings.Join(player.Inventory, ", ")
}

func who_description() string {
	var names []string
	for _, p := range players.Player {
		if p.Zone == player.Zone {
			names = append(names, p.Name)
		}
	}
	return "Room players: " + strings.Join(names, ", ")
}

func set_take() {
	inventory.RemoveAll()

	roomIdx := -1
	for i, r := range rooms.Rooms {
		if r.Name == player.Zone {
			roomIdx = i
		}
	}

	if roomIdx != -1 {
		for _, item := range rooms.Rooms[roomIdx].Items {
			inventory.Add(widget.NewButton(item, func() {
				player.Inventory = append(player.Inventory, item)
				items := rooms.Rooms[roomIdx].Items
				for j, it := range items {
					if it == item {
						rooms.Rooms[roomIdx].Items = append(items[:j], items[j+1:]...)
						break
					}
				}
				log_text.SetText("OK taken=" + item)
				set_inventory()
			}))
		}
	}
	inventory.Add(widget.NewButton("RETURN", func() {
		set_inventory()
	}))
	inventory.Refresh()
}

func set_drop() {
	inventory.RemoveAll()

	for _, item := range player.Inventory {
		inventory.Add(widget.NewButton(item, func() {
			for j, it := range player.Inventory {
				if it == item {
					player.Inventory = append(player.Inventory[:j], player.Inventory[j+1:]...)
					break
				}
			}
			for i, r := range rooms.Rooms {
				if r.Name == player.Zone {
					rooms.Rooms[i].Items = append(rooms.Rooms[i].Items, item)
				}
			}
			log_text.SetText("OK dropped=" + item)
			set_inventory()
		}))
	}
	inventory.Add(widget.NewButton("RETURN", func() {
		set_inventory()
	}))
	inventory.Refresh()
}
