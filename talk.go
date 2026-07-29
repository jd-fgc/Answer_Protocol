package main


func TALK(player_zone string, zones Rooms, npc_name string) string {
	var zone Room
	for _, val := range zones.Rooms {
		if val.Name == player_zone {
			zone = val
		}
	}
	if zone.Name == "" {
		return "ERR 404 NPC_NOT_FOUND"
	}
	for _, val := range zone.NPC {
		if val.Name == npc_name {
			return "OK " + val.Talk
		}
	}
	return "ERR 404 NPC_NOT_FOUND"
}
