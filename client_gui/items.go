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
	lines, err := netClient.SendCommand("INVENTORY")
	if err != nil {
		return "Erreur INVENTORY " + err.Error()
	}
	return lines[0]
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
		look, err := doLook()
		if err != nil {
			log_text.SetText("Erreur TAKE: " + err.Error())
		}
		rooms.Rooms[roomIdx].Items = look.Items
		for _, item := range rooms.Rooms[roomIdx].Items {
			inventory.Add(widget.NewButton(item, func() {
				lines, err := netClient.SendCommand("TAKE " + item)
				if err != nil || len(lines) == 0 {
					log_text.SetText("Erreur réseau lors du TAKE")
					return
				}
				log_text.SetText(strings.TrimSpace(lines[0]))
				player.Inventory = append(player.Inventory, item)
				refreshRoomState(false)
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

	for i, item := range player.Inventory {
		inventory.Add(widget.NewButton(item, func() {
			lines, err := netClient.SendCommand("DROP " + item)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du DROP")
				return
			}
			log_text.SetText(strings.TrimSpace(lines[0]))
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			refreshRoomState(false)
			set_inventory()
		}))
	}
	inventory.Add(widget.NewButton("RETURN", func() {
		set_inventory()
	}))
	inventory.Refresh()
}
