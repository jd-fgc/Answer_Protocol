package main

import "encoding/json"

func handleTake(From string, args []string) string {
	session, ok := sessions[From]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR missing item\n"
	}

	res := TAKE(session.Player, args[0])
	return res + "\n"
}

func TAKE(usr *Player, item_name string) string {
	for i, val := range rooms.Rooms {
		if val.Name == usr.Zone {
			for j, item := range val.Items {
				if item == item_name {
					usr.Inventory = append(usr.Inventory, item)
					rooms.Rooms[i].Items = append(rooms.Rooms[i].Items[:j], rooms.Rooms[i].Items[j+1:]...)
					update_players(*usr)
					return "OK taken=" + item_name
				}
			}
		}
	}
	return "ERR 404 ITEM_NOT_FOUND"
}

func handleDrop(From string, args []string) string {
	session, ok := sessions[From]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR missing item\n"
	}

	res := DROP(session.Player, args[0])
	return res + "\n"
}

func DROP(usr *Player, item_name string) string {
	for i, val := range rooms.Rooms {
		if val.Name == usr.Zone {
			for j, item := range usr.Inventory {
				if item == item_name {
					usr.Inventory = append(usr.Inventory[:j], usr.Inventory[j+1:]...)
					rooms.Rooms[i].Items = append(rooms.Rooms[i].Items, item_name)
					update_players(*usr)
					return "OK dropped=" + item_name
				}
			}
		}
	}
	return "ERR 404 ITEM_NOT_IN_INVENTORY"
}

func handleInventory(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}

	data, err := json.Marshal(session.Player.Inventory)
	if err != nil {
		return "ERR " + err.Error() + "\n"
	}
	return string(data) + "\n"
}

func INVENTORY(usr Player) (string, []string) {
	return "OK", usr.Inventory
}

// func USE(usr Player, zones Rooms, online_players Players, item_name string) error {

// }
