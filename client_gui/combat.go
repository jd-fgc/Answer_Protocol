package main

import (
	"fyne.io/fyne/v2/widget"
)

func set_attack() {
	refreshRoomState(false)
	combat.RemoveAll()
	for i := range rooms.Rooms {
		if rooms.Rooms[i].Name == player.Zone {
			rooms.Rooms[i].NPC = currentRoomNPCs
			room := rooms.Rooms[i]
			for j, npc := range room.NPC {
				i, j, npc := i, j, npc
				if npc.Ennemy == true {
					combat.Add(widget.NewButton(npc.Name, func() {
						lines, err := netClient.SendCommand("ATTACK " + npc.Name)
						if err != nil {
							return
						}
						if len(lines[0]) == 0 || len(lines) == 0 {
							return
						}
						rooms.Rooms[i].NPC[j].HP -= int(player.ATK*player.DMG_Mult)
						if rooms.Rooms[i].NPC[j].HP <= 0 {
							rooms.Rooms[i].NPC = append(rooms.Rooms[i].NPC[:j], rooms.Rooms[i].NPC[j+1:]...)
						}
						log_text.SetText(lines[0])
						refreshRoomState(false)
						set_log()
						set_combat()
					}))
				}
			}
		}
	}
	combat.Add(widget.NewButton("RETURN", func() {
		set_combat()
	}))
	combat.Refresh()
}