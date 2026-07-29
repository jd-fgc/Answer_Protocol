package main


func TAKE(usr *Player, zones *Rooms, online_players *Players, item_name string) string {
	for i, val := range zones.Rooms {
		if val.Name == usr.Zone {
			for j, item := range val.Items {
				if item == item_name {
					usr.Inventory = append(usr.Inventory, item)
					zones.Rooms[i].Items = append(zones.Rooms[i].Items[:j], zones.Rooms[i].Items[j+1:]...)
					update_players(*usr, online_players)
					return "OK taken=" + item_name
				}
			}
		}
	}
	return "ERR 404 ITEM_NOT_FOUND"
}

func DROP(usr *Player, zones *Rooms, online_players *Players, item_name string) string {
	for i, val := range zones.Rooms {
		if val.Name == usr.Zone {
			for j, item := range usr.Inventory {
				if item == item_name {
					usr.Inventory = append(usr.Inventory[:j], usr.Inventory[j+1:]...)
					zones.Rooms[i].Items = append(zones.Rooms[i].Items, item_name)
					update_players(*usr, online_players)
					return "OK dropped=" + item_name
				}
			}
		}
	}
	return "ERR 404 ITEM_NOT_IN_INVENTORY"
}


func INVENTORY(usr Player) (string, []string) {
	return "OK", usr.Inventory
}

// func USE(usr Player, zones Rooms, online_players Players, item_name string) error {
	
// }
