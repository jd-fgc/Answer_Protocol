package main

import (
	"fyne.io/fyne/v2/widget"
)

func set_quests() {
	actions.RemoveAll()

	for _, room := range rooms.Rooms {
		if room.Name == player.Zone {
			for _, val := range room.NPC {
				q := val.Quest
				for i, quest := range quests.Quest {
					if quest.Name == q && player.Zone == val.Zone && quest.taken == false {
						btn := widget.NewButton(quest.Name, func() {
							line, err := netClient.SendCommand("QUEST " + val.Name)
							if err != nil {
								return
							}
							log_text.SetText(line[0])
							quests.Quest[i].taken = true
							set_log()
							set_actions()
						})
						actions.Add(btn)
					}
				}
			}
		}
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}
