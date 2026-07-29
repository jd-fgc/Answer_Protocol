package main

import (
	"fmt"
)

func print_available_zone(zone Room) {
	fmt.Println("North: " + zone.North)
	fmt.Println("East: " + zone.East)
	fmt.Println("South: " + zone.South)
	fmt.Println("West: " + zone.West)
}

func check_next_zone(zone Room, zone_name string) bool {
	if zone.North == zone_name {
		return true
	}
	if zone.South == zone_name {
		return true
	}
	if zone.East == zone_name {
		return true
	}
	if zone.West == zone_name {
		return true
	}
	return false
}

func handleMove(from string, args []string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR missing direction\n"
	}

	res := MOVE(session.Player, rooms, args[0])
	return res + "\n"
}

func MOVE(player *Player, zones Rooms, zone_name string) string {
	for _, val := range zones.Rooms {
		if val.Name == player.Zone {
			if check_next_zone(val, zone_name) {
				player.Zone = zone_name
				update_players(*player, &online_players)
				return "OK room=" + zone_name
			}
		}
	}
	return "ERR 301 NO_EXIT"
}
