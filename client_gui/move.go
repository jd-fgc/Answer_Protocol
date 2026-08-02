package main

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

func current_room() Room {
	for _, r := range rooms.Rooms {
		if r.Name == player.zone {
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
		dest := dest
		actions.Add(widget.NewButton(label, func() {
			answer.cmd = "MOVE " + dest
			player.zone = dest
			visitedRooms[dest] = true
			playerViews[0].Player.zone = dest
			playerViews[0].Refresh()
			refreshMiniMap()
			log_text.SetText(fmt.Sprintf("OK room=%s", dest))
			set_actions()
		}))
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}
