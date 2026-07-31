package main


import (
	"fyne.io/fyne/v2/widget"
)


func set_talk() {
	actions.RemoveAll()

	for _, val := range npc {
		n := val

		if n.Zone == player.zone {
			btn := widget.NewButton(n.Name, func() {
				answer.cmd = "TALK " + n.Name
				log_text.Text = n.Talk
				set_actions()
				set_log()
			})
			actions.Add(btn)
		}
	}
	actions.Add(widget.NewButton("RETURN", func() {
				set_actions()
	}))
	actions.Refresh()
}