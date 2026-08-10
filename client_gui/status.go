package main

import (
	"encoding/json"
	"strings"
)

type playerStatus struct {
	CURRENT_HP int64   `json:"hp"`
	MaxHP      int64   `json:"max_hp"`
	ATK        float64 `json:"atk"`
	Coin       int64   `json:"coin"`
	Bomb       int64   `json:"bomb"`
	Status     string  `json:"status"`
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
	player.CURRENT_HP = int(s.CURRENT_HP)
	player.MAX_HP = int(s.MaxHP)
	player.ATK = float32(s.ATK)
	player.COIN = int(s.Coin)
	player.BOMB = int(s.Bomb)

	if pv := findPlayerViewSelf(); pv != nil {
		pv.Refresh()
	}
}

func findPlayerViewSelf() *PlayerView {
	if selfPlayerView != nil {
		return selfPlayerView
	}
	for _, pv := range playerViews {
		if pv.Player.Name == player.Name {
			return pv
		}
	}
	return nil
}