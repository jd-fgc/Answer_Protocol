package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
)


func load_quests(file_path string) (Quests, error) {
	var q Quests
	file, err := os.Open(file_path)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &q)
	if err != nil {
		return Quests{Quest: nil}, errors.New(err.Error())
	}
	return q, nil
}

func handleQuest(from string, args []string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR missing NPC name\n"
	}

	npc_name := ""
	for i, val := range args{
		npc_name += val
		if i < len(args) - 1 {
			npc_name += " "
		}
	}
	result := QUEST(session.Player, npc_name)
	return result + "\n"
}

func handleQuests(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	result := QUESTS()
	return result + "\n"
}

func handleCompleteQuest(from string, args []string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR No quest name provided\n"
	}

	quest_name := strings.Join(args, " ")
	result := COMPLETE_QUEST(session.Player, quest_name)
	return result + "\n"
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
							update_players(*usr)
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

func get_quest_reward(usr *Player, quest Quest) {
	switch quest.Reward {
	case "coin":
		usr.Coin += quest.Amount_reward
	case "bomb":
		usr.Bomb += quest.Amount_reward
	case "key":
		usr.Key += quest.Amount_reward
	default:
		for i := 0; i < quest.Amount_reward; i++ {
			usr.Inventory = append(usr.Inventory, quest.Reward)
		} 
	}
}

func COMPLETE_QUEST(usr *Player, quest_name string) string {
	for i, val := range quests.Quest {
		if val.Name == quest_name && val.Completed == true && val.finished == false {
			get_quest_reward(usr, val)
			update_players(*usr)
			quests.Quest[i].finished = true
			return "OK Completed " + quest_name
		}
	}
	return "ERR 405 QUEST_NOT_FOUND_OR_NOT_COMPLETED"
}