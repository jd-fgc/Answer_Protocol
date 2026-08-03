package main

func handleTalk(from string, args []string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR missing npc name\n"
	}

	var npc_name string
	for i, w := range args {
		npc_name += w
		if len(args) > i+1 {
			npc_name += " "
		}
	}
	result := TALK(session.Player.Zone, npc_name)
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
