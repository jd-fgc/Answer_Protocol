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

func check_dir(zone Room, dir string) bool {
	if dir == "north" && zone.North != "" {
		return true
	}
	if dir == "south" && zone.South != "" {
		return true
	}
	if dir == "east" && zone.East != "" {
		return true
	}
	if dir == "west" && zone.West != "" {
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

func MOVE(player *Player, zone_name string) string {
	for _, val := range rooms.Rooms {
		if val.Name == player.Zone {
			if check_next_zone(val, zone_name) {
				player.Zone = zone_name
				update_players(*player)
				return "OK room=" + zone_name
			} else if check_dir(val, zone_name) {
				switch zone_name {
				case "north":
					player.Zone = val.North
				case "south":
					player.Zone = val.South
				case "east":
					player.Zone = val.East
				case "west":
					player.Zone = val.West
				}
				update_players(*player)
				return "OK room=" + player.Zone
			}
		}
	}
	return "ERR 301 NO_EXIT"
}
