package main

import "encoding/json"

type Who_Struct struct {
	Room   []string `json:"room"`
	Server int      `json:"server"`
}

func WHO(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "Err not connected\n"
	}

	var roomPlayers []string
	for _, p := range online_players.Players {
		if p.Zone == session.Player.Zone {
			roomPlayers = append(roomPlayers, p.Name)
		}
	}

	who := Who_Struct{
		Room:   roomPlayers,
		Server: len(online_players.Players),
	}

	data, err := json.Marshal(who)
	if err != nil {
		return "ERR " + err.Error() + "\n"
	}
	return string(data) + "\n"
}

// func CHAT(from string) string {

// }
