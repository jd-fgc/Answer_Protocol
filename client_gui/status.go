package main

import (
	"encoding/json"
	"strings"
)

type playerStatus struct {
	HP     int64  `json:"hp"`
	MaxHP  int64  `json:"max_hp"`
	Status string `json:"status"`
}

func doStatus() (*playerStatus, error) {
	lines, err := netClient.SendCommand("STATUS")
	if err != nil || len(lines) == 0 {
		return nil, err
	}

	line := strings.TrimSpace(lines[0])
	line = strings.TrimPrefix(line, "OK ")

	var s playerStatus
	if err := json.Unmarshal([]byte(line), &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func refreshPlayerStatus() {
	s, err := doStatus()
	if err != nil {
		return
	}
	player.HP = int(s.HP)

	if selfPlayerView != nil {
		selfPlayerView.Refresh()
	}
}

func findPlayerViewSelf() *PlayerView {
	for _, pv := range playerViews {
		if pv.Player.Name == player.Name {
			return pv
		}
	}
	return nil
}
