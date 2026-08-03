package main

import (
	"strings"

	"fyne.io/fyne/v2/widget"
)

func current_room() Room {
	for _, r := range rooms.Rooms {
		if r.Name == player.Zone {
			return r
		}
	}
	return Room{}
}

func set_move() {
	actions.RemoveAll()

	r := current_room()
	dirs := map[string]string{
		"North": r.North,
		"South": r.South,
		"East":  r.East,
		"West":  r.West,
	}

	for label, dest := range dirs {
		if dest == "" {
			continue
		}
		label, dest := label, dest
		actions.Add(widget.NewButton(label, func() {
			lines, err := netClient.SendCommand("MOVE " + dest)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du MOVE")
				return
			}
			if !strings.HasPrefix(lines[0], "OK") {
				log_text.SetText(strings.TrimSpace(lines[0]))
				return
			}

			refreshRoomState()
			set_actions()
		}))
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}
