package main

import (
	// "errors"
	"encoding/json"
)


type Look_Struct struct {
	Room Room_Description `json:"room"`
	Players []string `json:"players"`
	Items []string `json:"items"`
	NPCs []string `json:"npcs"`
}


type Room_Description struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Exits Exits `json:"exits"`
}

type Exits struct {
	North string `json:"north,omitempty"`
	South string `json:"south,omitempty"`
	East string `json:"east,omitempty"`
	West string `json:"west,omitempty"`
}


func fill_desc(zone Room) Look_Struct {
	var look_struct Look_Struct
	var room_desc Room_Description
	var exits Exits

	exits.North = zone.North
	exits.South = zone.South
	exits.East = zone.East
	exits.West = zone.West

	room_desc.Id = "0" // zone.Id
	room_desc.Name = zone.Name
	room_desc.Description = "" // zone.Description
	room_desc.Exits = exits

	look_struct.Room = room_desc
	for _, val := range online_player.Players {
		if val.Zone == zone.Name {
			look_struct.Players = append(look_struct.Players, val.Name)
		}
	}
	look_struct.Items = zone.Items
	for _, val := range zone.NPC {
		look_struct.NPCs = append(look_struct.NPCs, val.Name)
	}

	return look_struct
}


func LOOK(player Player) (string, error) {
	for _, val := range rooms.Rooms {
		if val.Name == player.Zone {
			desc := fill_desc(val)
			data, err := json.Marshal(desc)
			return string(data), err
		}
	}
	data, err := json.Marshal("")
	return string(data), err
}
