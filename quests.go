package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)


func load_quests(file_path string) (Quests, error) {
	var quests []Quest
	file, err := os.Open(file_path)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &quests)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	return Quests{Quest: quests}, nil
}


func QUEST(usr *Player, npc_name string) (string) {
	for _, room := range rooms.Rooms {
		if room.Name == usr.Zone {
			for _, npc := range room.NPC {
				if npc.Name == npc_name {
					for k, qst := range quests.Quest {
						if qst.Name == npc.Quest && qst.taken == false {
							data, err := json.Marshal(quests.Quest[k])
							if err != nil {
								return err.Error()
							}
							quests.Quest[k].taken = true
							usr.Quests = append(usr.Quests, quests.Quest[k])
							return "OK " + string(data)
						}
					}
					return "ERR 406 NO_QUEST_AVAILABLE"
				}
			}
		}
	}
	return "ERR 404 NPC_NOT_FOUND"
}


func QUESTS() string {
	var active_quests []Quest
	for _, quest := range quests.Quest {
		if quest.taken == true {
			active_quests = append(active_quests, quest)
		}
	}
	data, err := json.Marshal(active_quests)
	if err != nil {
		return err.Error()
	}
	return "OK " + string(data)
}


// func COMPLETE_QUEST(usr *Player, quest_name string) string {
	
// }