package main

func handleTalk(from string, args []string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR missing npc name\n"
	}

	result := TALK(session.Player.Zone, args[0])
	return result + "\n"
}

func TALK(player_zone string, npc_name string) string {
	var zone Room
	for _, val := range rooms.Rooms {
		if val.Name == player_zone {
			zone = val
		}
	}
	if zone.Name == "" {
		return "ERR 404 NPC_NOT_FOUND"
	}
	for _, val := range zone.NPC {
		if val.Name == npc_name {
			return val.Talk
		}
	}
	return "ERR 404 NPC_NOT_FOUND"
}
