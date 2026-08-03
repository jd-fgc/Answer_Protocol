package main

import "encoding/json"

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
