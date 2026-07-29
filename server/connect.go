package main

import "encoding/json"

type Session struct {
	Username string
	Player   *Player
}

var sessions = make(map[string]*Session)

func handleConnect(from string, args []string) string {
	if len(args) == 0 {
		return "ERR missing name\n"
	}
	username := args[0]

	for _, s := range sessions {
		if s.Username == username {
			return "ERR name already taken\n"
		}
	}

	sessions[from] = &Session{Username: username}

	var names []string
	for _, c := range playable_characters.Players {
		names = append(names, c.Name)
	}

	data, err := json.Marshal(names)
	if err != nil {
		return "ERR " + err.Error() + "\n"
	}

	return "OK connected\n" + string(data) + "\n"
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
	template.Name = session.Username

	session.Player = &template
	online_players.Players = append(online_players.Players, template)

	return "OK character=" + args[0] + "\n"
}
