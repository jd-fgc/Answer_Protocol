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

func get_room_dir(dest string, room Room) string {
	switch dest {
	case room.North:
		return "North"
	case room.South:
		return "South"
	case room.East:
		return "East"
	case room.West:
		return "west"
	default:
		return "ERROR"
	}
}

func set_move() {
	actions.RemoveAll()

	r := current_room()
	dirs := []string{
		r.North,
		r.South,
		r.East,
		r.West,
	}

	for _, dest := range dirs {
		if dest == "" {
			continue
		}
		label, dest := get_room_dir(dest, r), dest
		actions.Add(widget.NewButton(label, func() {
			lines, err := netClient.SendCommand("MOVE " + dest)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du MOVE")
				return
			}
			if strings.HasPrefix(lines[0], "OK") {
				log_text.SetText(strings.TrimSpace(lines[0]))
			}

			refreshRoomState(false)
			set_actions()
		}))
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}
