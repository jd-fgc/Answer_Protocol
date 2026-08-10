package main

import (
	"strings"

	"fyne.io/fyne/v2/widget"
)

func set_talk() {
	actions.RemoveAll()

	for _, name := range currentRoomNPCsName {
		name := name
		actions.Add(widget.NewButton(name, func() {
			lines, err := netClient.SendCommand("TALK " + name)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du TALK")
				return
			}
			log_text.SetText(strings.TrimSpace(lines[0]))
			set_actions()
		}))
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}
