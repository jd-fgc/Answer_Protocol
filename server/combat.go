package main

import "encoding/json"

type ATTACK_TURN struct {
	Attacker_hp int    `json:"attacker_hp"`
	Target_hp   int    `json:"target_hp"`
	Damage      int    `json:"damage"`
	Status      string `json:"status"`
}

type PLAYER_STATUS struct {
	HP     int    `json:"hp"`
	MAX_HP int    `json:"max_hp"`
	STATUS string `json:"status"`
}

func fill_attack_info(usr Player, npc NPC, damage int) string {
	var turn ATTACK_TURN
	turn.Attacker_hp = usr.CURRENT_HP
	turn.Target_hp = npc.HP
	turn.Damage = damage
	turn.Status = "combat"

	data, err := json.Marshal(turn)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func handleAttack(from string, args []string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR missing Ennemy\n"
	}

	enemy_name := ""
	for i, val := range args {
		enemy_name += val
		if i < len(args)-1 {
			enemy_name += " "
		}
	}
	result := ATTACK(*session.Player, enemy_name)
	return result + "\n"
}

func handleDefend(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	result := DEFEND(session.Player)
	return result + "\n"
}

func handleStatus(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	result := STATUS(*session.Player)
	return result + "\n"
}

func ATTACK(usr Player, npc_name string) string {
	for i, val := range rooms.Rooms {
		for j, npc := range val.NPC {
			if npc.Name == npc_name && npc.Ennemy == true {
				npc.HP = npc.HP - int(usr.ATK*usr.DMG_Mult)
				rooms.Rooms[i].NPC[j] = npc
				if npc.HP <= 0 {
					npc.HP = 0
					rooms.Rooms[i].NPC = append(rooms.Rooms[i].NPC[:j], rooms.Rooms[i].NPC[j+1:]...)
				}
				return fill_attack_info(usr, npc, int(usr.ATK*usr.DMG_Mult))
			}
			if npc.Name == npc_name && npc.Ennemy == false {
				return "ERR 405 NPC_NOT_HOSTILE"
			}
		}
	}
	return "ERR 404 NPC_NOT_FOUND"
}

func DEFEND(usr *Player) string {
	usr.Damage_Reduction = 0.25
	update_players(*usr)
	return "OK " + usr.Name + " ready to defend"
}

func STATUS(usr Player) string {
	var status PLAYER_STATUS
	status.HP = usr.CURRENT_HP
	status.MAX_HP = usr.MAX_HP
	if usr.CURRENT_HP >= int(usr.MAX_HP*80/100) {
		status.STATUS = "healthy !"
	} else if usr.CURRENT_HP >= int(usr.MAX_HP*50/100) {
		status.STATUS = "could be better."
	} else if usr.CURRENT_HP >= int(usr.MAX_HP*30/100) {
		status.STATUS = "death seems inevitable..."
	} else {
		status.STATUS = "Either run or heal but for the love of God, don't die."
	}
	data, err := json.Marshal(status)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
