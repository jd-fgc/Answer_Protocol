package main

import "encoding/json"

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
			return "ERR name already taken\n"
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

	logInfo("character_chosen", map[string]interface{}{"from": from, "username": session.Username, "character": args[0]})
	return "OK character=" + args[0] + "\n"
}

type WhoPlayerInfo struct {
	Name      string `json:"name"`
	Character string `json:"character"`
}

type Who_Struct struct {
	Room   []WhoPlayerInfo `json:"room"`
	Server int             `json:"server"`
}

func handleWho(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}

	roomPlayers := []WhoPlayerInfo{}
	for _, p := range online_players.Players {
		if p.Zone == session.Player.Zone {
			roomPlayers = append(roomPlayers, WhoPlayerInfo{Name: p.Name, Character: p.Character})
		}
	}

	who := Who_Struct{Room: roomPlayers, Server: len(online_players.Players)}
	data, err := json.Marshal(who)
	if err != nil {
		return "ERR " + err.Error() + "\n"
	}
	return string(data) + "\n"
}
