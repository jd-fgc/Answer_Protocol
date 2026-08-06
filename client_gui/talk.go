package main

import (
	"fyne.io/fyne/v2/widget"
)

func set_talk() {
	actions.RemoveAll()

	for _, room := range rooms.Rooms {
		if room.Name == player.Zone {
			for _, npc := range room.NPC {
				btn := widget.NewButton(npc.Name, func() {
					lines, err := netClient.SendCommand("TALK " + npc.Name)
					if err != nil {
						return
					}
					log_text.SetText(lines[0])
					set_actions()
					set_log()
				})
				actions.Add(btn)
			}
		}
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}
