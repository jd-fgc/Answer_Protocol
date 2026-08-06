package main

import (
	// "errors"
	"encoding/json"
)

type Room_Description struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Exits       Exits  `json:"exits"`
}

type Exits struct {
	North string `json:"north,omitempty"`
	South string `json:"south,omitempty"`
	East  string `json:"east,omitempty"`
	West  string `json:"west,omitempty"`
}

type PlayerInfo struct {
	Name      string `json:"name"`
	Character string `json:"character"`
}

type Look_Struct struct {
	Room    Room_Description `json:"room"`
	Players []PlayerInfo     `json:"players"`
	Items   []ItemInstance   `json:"items"`
	NPCs    []string         `json:"npcs"`
}

func fill_desc(players Players, zone Room) Look_Struct {
	var look_struct Look_Struct
	var room_desc Room_Description
	var exits Exits

	exits.North = zone.North
	exits.South = zone.South
	exits.East = zone.East
	exits.West = zone.West

	room_desc.Id = "room." + slugify(zone.Name)
	room_desc.Name = zone.Name
	room_desc.Description = ""
	room_desc.Exits = exits

	look_struct.Room = room_desc
	for _, val := range players.Players {
		if val.Zone == zone.Name {
			look_struct.Players = append(look_struct.Players, PlayerInfo{Name: val.Name, Character: val.Character})
		}
	}
	look_struct.Items = zone.Items
	for _, val := range zone.NPC {
		look_struct.NPCs = append(look_struct.NPCs, val.Name)
	}

	return look_struct
}

func handleLook(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}

	result, err := LOOK(*session.Player, rooms, online_players)
	if err != nil {
		return "ERR " + err.Error() + "\n"
	}
	return "OK " + result + "\n"
}

func LOOK(player Player, zones Rooms, players Players) (string, error) {
	for _, val := range zones.Rooms {
		if val.Name == player.Zone {
			desc := fill_desc(players, val)
			data, err := json.Marshal(desc)
			return string(data), err
		}
	}
	data, err := json.Marshal("")
	return string(data), err
}
