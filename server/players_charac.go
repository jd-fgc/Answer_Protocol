package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func update_players(player Player) error {
	for i, val := range online_players.Players {
		if val.Name == player.Name {
			online_players.Players[i] = player
			return nil
		}
	}
	return errors.New("Player not found")
}

func read_character_file(filepath string, start_zone string) (Players, error) {
	var charac []Player
	file, err := os.Open(filepath)
	if err != nil {
		return Players{Players: nil}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Players{Players: nil}, errors.New(err.Error())
	}
	err = json.Unmarshal(byteValue, &charac)
	if err != nil {
		return Players{Players: nil}, errors.New(err.Error())
	}
	for i := range charac {
		charac[i].Zone = start_zone
	}
	return Players{Players: charac}, nil
}

type rawRoom struct {
	Name  string   `json:"name"`
	NPC   []NPC    `json:"NPC"`
	Items []string `json:"item"`
	North string   `json:"North"`
	South string   `json:"South"`
	East  string   `json:"East"`
	West  string   `json:"West"`
}

type rawRooms struct {
	Rooms []rawRoom `json:"world"`
}

func slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "_")
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func load_map(filepath string) (Rooms, error) {
	var raw rawRooms
	file, err := os.Open(filepath)
	if err != nil {
		return Rooms{}, errors.New(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Rooms{}, errors.New(err.Error())
	}
	if err := json.Unmarshal(byteValue, &raw); err != nil {
		return Rooms{}, errors.New(err.Error())
	}

	var rooms Rooms
	counter := map[string]int{}

	for _, rr := range raw.Rooms {
		room := Room{Name: rr.Name, NPC: rr.NPC, North: rr.North, South: rr.South, East: rr.East, West: rr.West}
		for _, name := range rr.Items {
			slug := slugify(name)
			counter[slug]++
			id := "item." + slug
			if counter[slug] > 1 {
				id = fmt.Sprintf("item.%s-%d", slug, counter[slug])
			}
			room.Items = append(room.Items, ItemInstance{ID: id, Name: name})
			itemDisplayName[id] = name
		}
		rooms.Rooms = append(rooms.Rooms, room)
	}
	for i, rq := range quests.Quest {
		switch rq.Reward {
		case "coin":
		case "bomb":
		case "key":
		default:
			slug := slugify(rq.Reward)
			counter[slug]++
			id := "item." + slug
			if counter[slug] > 1 {
				id = fmt.Sprintf("item.%s-%d", slug, counter[slug])
			}
			itemDisplayName[id] = rq.Reward
			quests.Quest[i].Reward = id
		}
	}

	return rooms, nil
}

func character_selection(characters Players, user_choice string) (Player, error) {
	for _, val := range characters.Players {
		if val.Name == user_choice {
			return val, nil
		}
	}
	return Player{}, errors.New("No player found with this name")
}
