package main

import (
	"encoding/json"
	"strings"
)

func findItemInRoom(roomIdx int, query string) (int, ItemInstance, bool) {
	for i, it := range rooms.Rooms[roomIdx].Items {
		if it.ID == query || it.Name == query {
			return i, it, true
		}
	}
	return -1, ItemInstance{}, false
}

func handleTake(From string, args []string) string {
	session, ok := sessions[From]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		return "ERR missing item\n"
	}
	query := strings.Join(args, " ")
	res := TAKE(session.Player, query)
	return res + "\n"
}

func TAKE(usr *Player, query string) string {
	for i, val := range rooms.Rooms {
		if val.Name == usr.Zone {
			idx, it, found := findItemInRoom(i, query)
			if !found {
				return "ERR 404 ITEM_NOT_FOUND"
			}
			usr.Inventory = append(usr.Inventory, it.ID)
			rooms.Rooms[i].Items = append(rooms.Rooms[i].Items[:idx], rooms.Rooms[i].Items[idx+1:]...)
			update_players(*usr)
			logInfo("item_taken", map[string]interface{}{"player": usr.Name, "item_id": it.ID, "room": usr.Zone})
			return "OK taken=" + it.ID
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
	query := strings.Join(args, " ")
	res := DROP(session.Player, query)
	return res + "\n"
}

func DROP(usr *Player, query string) string {
	for i, id := range usr.Inventory {
		if id == query || itemDisplayName[id] == query {
			name := itemDisplayName[id]
			usr.Inventory = append(usr.Inventory[:i], usr.Inventory[i+1:]...)
			for ri, r := range rooms.Rooms {
				if r.Name == usr.Zone {
					rooms.Rooms[ri].Items = append(rooms.Rooms[ri].Items, ItemInstance{ID: id, Name: name})
				}
			}
			update_players(*usr)
			logInfo("item_dropped", map[string]interface{}{"player": usr.Name, "item_id": id, "room": usr.Zone})
			return "OK dropped=" + id
		}
	}
	return "ERR 404 ITEM_NOT_IN_INVENTORY"
}

func handleInventory(from string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}

	var items []ItemInstance
	for _, id := range session.Player.Inventory {
		items = append(items, ItemInstance{ID: id, Name: itemDisplayName[id]})
	}

	data, err := json.Marshal(items)
	if err != nil {
		return "ERR " + err.Error() + "\n"
	}
	return string(data) + "\n"
}

// func USE(usr Player, zones Rooms, online_players Players, item_name string) error {

// }
