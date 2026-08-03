package main

import (
	"fyne.io/fyne/v2/widget"
)

func get_quest_info(quest_name string) string {
	for _, val := range quests.Quest {
		if val.Name == quest_name {
			return val.Description
		}
	}
	return ""
}

func set_quests() {
	actions.RemoveAll()

	for _, val := range npc {
		q := val.Quest

		if q != "" && player.Zone == val.Zone {
			btn := widget.NewButton(q, func() {
				answer.cmd = "QUEST " + val.Name
				log_text.Text = get_quest_info(q)
				set_log()
				set_actions()
			})
			actions.Add(btn)
		}
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}
