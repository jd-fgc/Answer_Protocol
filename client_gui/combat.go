package main

import "fyne.io/fyne/v2/widget"

func set_attack() {
	combat.RemoveAll()
	for i, room := range rooms.Rooms {
		if room.Name == player.Zone {
			for j, npc := range room.NPC {
				if npc.Ennemy {
					i, j, npc := i, j, npc
					combat.Add(widget.NewButton(npc.Name, func() {
						lines, err := netClient.SendCommand("ATTACK " + npc.Name)
						if err != nil {
							return
						}
						if len(lines) == 0 {
							return
						}
						rooms.Rooms[i].NPC[j].HP -= int(player.ATK * player.DMG_Mult)
						if rooms.Rooms[i].NPC[j].HP <= 0 {
							rooms.Rooms[i].NPC = append(rooms.Rooms[i].NPC[:j], rooms.Rooms[i].NPC[j+1:]...)
						}
						log_text.SetText(lines[0])
						refreshPlayerStatus()
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
