package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type playerInfo struct {
	Name      string `json:"name"`
	Character string `json:"character"`
}

type lookResponse struct {
	Room    lookRoomDesc   `json:"room"`
	Players []playerInfo   `json:"players"`
	Items   []itemInstance `json:"items"`
	NPCs    []string       `json:"npcs"`
}

type lookExits struct {
	North string `json:"north,omitempty"`
	South string `json:"south,omitempty"`
	East  string `json:"east,omitempty"`
	West  string `json:"west,omitempty"`
}

type lookRoomDesc struct {
	Name  string    `json:"name"`
	Exits lookExits `json:"exits"`
}

type itemInstance struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func doLook() (*lookResponse, error) {
	lines, err := netClient.SendCommand("LOOK")
	if err != nil || len(lines) == 0 {
		return nil, err
	}

	line := strings.TrimSpace(lines[0])
	line = strings.TrimPrefix(line, "OK ")

	var resp lookResponse
	if err := json.Unmarshal([]byte(line), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func updateKnownRoom(name string, exits lookExits) {
	for i, r := range rooms_map.Rooms {
		if r.Name == name {
			rooms_map.Rooms[i].North = exits.North
			rooms_map.Rooms[i].South = exits.South
			rooms_map.Rooms[i].East = exits.East
			rooms_map.Rooms[i].West = exits.West
			return
		}
	}
	rooms_map.Rooms = append(rooms_map.Rooms, Room{Name: name, North: exits.North, South: exits.South, East: exits.East, West: exits.West})
}

var currentRoomItems []itemInstance
var currentRoomNPCs []string

func refreshRoomState(is_look bool) {
	resp, err := doLook()
	if err != nil {
		log_text.SetText("Erreur LOOK: " + err.Error())
		return
	}

	player.Zone = resp.Room.Name
	visitedRooms[resp.Room.Name] = true
	updateKnownRoom(resp.Room.Name, resp.Room.Exits)
	currentRoomItems = resp.Items
	currentRoomNPCs = resp.NPCs

	refreshMiniMap()
	refreshRoomPlayers(leftColumnGlobal, resp.Players)

	if is_look {
		var itemNames []string
		for _, it := range resp.Items {
			itemNames = append(itemNames, it.Name)
		}

		log_text.SetText(fmt.Sprintf("Room: %s\nItems: %s\nNPCs: %s",
			resp.Room.Name, strings.Join(itemNames, ", "), strings.Join(resp.NPCs, ", ")))
	}
}
