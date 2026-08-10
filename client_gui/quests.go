package main

import (
	"fyne.io/fyne/v2/widget"
)

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
							line, err := netClient.SendCommand("QUEST " + val.Name)
							if err != nil {
								return
							}
							player.Quests = append(player.Quests, quest)
							log_text.SetText(line[0])
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
		// for i := 0; i < quest.Amount_reward; i++ {
		// 	player.Inventory = append(player.Inventory, quest.Reward)
		// } 
	}
}