package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)


func load_quests(file_path string) (QUESTS, error) {
	var quests []Quest
	file, err := os.Open(file_path)
	if err != nil {
		return QUESTS{Quest: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return QUESTS{Quest: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &quests)
	if err != nil {
		return QUESTS{Quest: nil}, errors.New(err.Error())
	}
	return QUESTS{Quest: quests}, nil
}


func QUEST(usr Player, rooms Rooms, quests QUESTS, npc_name string) (string) {
	for _, room := range rooms.Rooms {
		if room.Name == usr.Zone {
			
		}
	}
}
