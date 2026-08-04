package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

type lookResponse struct {
	Room    lookRoomDesc `json:"room"`
	Players []string     `json:"players"`
	Items   []string     `json:"items"`
	NPCs    []string     `json:"npcs"`
}

func doLook() (*lookResponse, error) {
	lines, err := netClient.SendCommand("LOOK")
	if err != nil || len(lines) == 0 {
		return nil, err
	}
	var resp lookResponse
	if err := json.Unmarshal([]byte(lines[0]), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func updateKnownRoom(name string, exits lookExits) {
	for i, r := range rooms.Rooms {
		if r.Name == name {
			rooms.Rooms[i].North = exits.North
			rooms.Rooms[i].South = exits.South
			rooms.Rooms[i].East = exits.East
			rooms.Rooms[i].West = exits.West
			return
		}
	}
	rooms.Rooms = append(rooms.Rooms, Room{Name: name, North: exits.North, South: exits.South, East: exits.East, West: exits.West})
}

var currentRoomItems []string
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
	refreshRoomPlayers(leftColumnGlobal)

	if is_look {
		log_text.SetText(fmt.Sprintf("Room: %s\nItems: %s\nNPCs: %s",
			resp.Room.Name, strings.Join(resp.Items, ", "), strings.Join(resp.NPCs, ", ")))
	}
}
