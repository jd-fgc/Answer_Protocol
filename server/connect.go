package main

import (
	"encoding/json"
	"fmt"
)

type Session struct {
	Username string
	Player   *Player
	Group    string
	Events   chan string
}

var sessions = make(map[string]*Session)

func handleConnect(from string, args []string) string {
	if len(args) == 0 {
		return "ERR missing name\n"
	}
	username := args[0]

	if len(online_players.Players) >= 4 {
		return "ERR server full\n"
	}

	for _, s := range sessions {
		if s.Username == username {
			return "ERR 201 NAME_IN_USE\n"
		}
	}

	sessions[from] = &Session{Username: username, Events: make(chan string, 20)}

	var names []string
	for _, c := range playable_characters.Players {
		names = append(names, c.Name)
	}

	data, err := json.Marshal(names)
	if err != nil {
		return "ERR " + err.Error() + "\n"
	}

	logInfo("player_connected", map[string]interface{}{"from": from, "username": username})
	return "OK connected\n" + string(data) + "\n"
}

func handleDisconnect(from string) {
	session, ok := sessions[from]
	if !ok {
		return
	}
	if session.Player != nil {
		for i, p := range online_players.Players {
			if p.Name == session.Player.Name {
				online_players.Players = append(online_players.Players[:i], online_players.Players[i+1:]...)
				break
			}
		}
	}
	delete(sessions, from)
}

func handleQuit(from string) string {
	session, ok := sessions[from]
	if !ok {
		return "ERR not connected\n"
	}

	if session.Player != nil {
		for i, p := range online_players.Players {
			if p.Name == session.Player.Name {
				online_players.Players = append(online_players.Players[:i], online_players.Players[i+1:]...)
				broadcastGlobal(fmt.Sprintf("EVT STATS PLAYERS players=%d\n", len(online_players.Players)))
				break
			}
		}
	}

	delete(sessions, from)
	logInfo("player_quit", map[string]interface{}{"from": from})
	return "OK bye\n"
}

func handleChoose(from string, args []string) string {
	if len(args) == 0 {
		return "ERR missing character name\n"
	}

	session, ok := sessions[from]
	if !ok || session.Username == "" {
		return "ERR not connected\n"
	}
	if session.Player != nil {
		return "ERR character already chosen\n"
	}

	template, err := character_selection(playable_characters, args[0])
	if err != nil {
		return "ERR character not found\n"
	}
	template.Character = args[0]
	template.Name = session.Username

	session.Player = &template
	online_players.Players = append(online_players.Players, template)
	broadcastGlobal(fmt.Sprintf("EVT STATS PLAYERS players=%d\n", len(online_players.Players)))

	logInfo("character_chosen", map[string]interface{}{"from": from, "username": session.Username, "character": args[0]})
	return "OK character=" + args[0] + "\n"
}

func handleWho(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	return fmt.Sprintf("OK players=%d\n", len(online_players.Players))
}
