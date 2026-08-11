package main

import (
	"encoding/json"
	"strings"
)

type playerStatus struct {
	CURRENT_HP int     `json:"hp"`
	MaxHP      int     `json:"max_hp"`
	ATK        float32 `json:"atk"`
	Coin       int     `json:"coin"`
	Bomb       int     `json:"bomb"`
	Key        int     `json:"key"`
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
	player.KEY = s.Key

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
