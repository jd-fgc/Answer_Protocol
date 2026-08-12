package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)


type questInfo struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}


func set_quests() {
	refreshRoomState(false)
	actions.RemoveAll()

	for _, room := range rooms.Rooms {
		if room.Name == player.Zone {
			for _, val := range room.NPC {
				q := val.Quest
				for i, quest := range quests.Quest {
					if quest.Name == q && player.Zone == val.Zone && quest.taken == false {
						btn := widget.NewButton(quest.Name, func() {
							lines, err := netClient.SendCommand("QUEST " + val.Name)
							if err != nil {
								return
							}
							player.Quests = append(player.Quests, quest)
							line := strings.Trim(lines[0], "OK ")
							var qInfo questInfo
							err = json.Unmarshal([]byte(line), &qInfo)
							if err != nil {
								return
							}
							log_text.SetText(fmt.Sprintf("Quest Name: %s. Quest Description: %s", quest.Name, quest.Description))
							quests.Quest[i].taken = true
							refreshRoomState(false)
							set_log()
							set_actions()
						})
						actions.Add(btn)
					}
				}
			}
		}
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}

func set_complete_quests() {
	actions.RemoveAll()
	update_rooms()
	check_quests()
	for i, quest := range player.Quests {
		if quest.Completed == true && quest.finished == false{
			actions.Add(widget.NewButton(quest.Name, func() {
				lines, err := netClient.SendCommand("COMPLETE_QUEST " + quest.Name)
				if err != nil {
					return
				}
				player.Quests[i].finished = true
				log_text.SetText(lines[0])
				get_quest_reward(quest)
				set_actions()
			}))
		}
	}
	actions.Add(widget.NewButton("RETURN", func() {
		set_actions()
	}))
	actions.Refresh()
}

func check_quests() {
	tmp_usr := player

	for i, quest := range tmp_usr.Quests {
		switch quest.Type {
		case "item":
			switch quest.Item_needed {
			case "coin":
				if tmp_usr.COIN - quest.Needed_amount >= 0 {
					player.Quests[i].Completed = true
				} else {
					player.Quests[i].Completed = false
				}
			case "bomb":
				if tmp_usr.BOMB - quest.Needed_amount >= 0 {
					player.Quests[i].Completed = true
				} else {
					player.Quests[i].Completed = false
				}
			case "key":
				if tmp_usr.KEY - quest.Needed_amount >= 0 {
					player.Quests[i].Completed = true
				} else {
					player.Quests[i].Completed = false
				}
			default:
			}
		case "ennemy":
			ennemy_alive := false
			for _, room := range rooms.Rooms {
				for _, npc := range room.NPC {
					if npc.Name == quest.Ennemy {
						ennemy_alive = true
					}
				}
			}
			if ennemy_alive == false {
				player.Quests[i].Completed = true
			}
		default:
		}
	}
}

func get_quest_reward(quest Quest) {
	switch quest.Reward {
	case "coin":
		player.COIN += quest.Amount_reward
	case "bomb":
		player.BOMB += quest.Amount_reward
	case "key":
		player.KEY += quest.Amount_reward
	default:
	}
}

func update_rooms() {
	data, err := netClient.SendCommand("REFRESH_ROOMS")
	if err != nil {
		return
	}
	var n_rooms Rooms
	err = json.Unmarshal([]byte(data[0]), &n_rooms)
	if err != nil {
		return
	}
	rooms = n_rooms
}