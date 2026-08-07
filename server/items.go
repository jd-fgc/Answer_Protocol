package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"
)

func read_items_file(filepath string) (*ItemsFile, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var fileData ItemsFile
	err = json.Unmarshal(byteValue, &fileData)
	if err != nil {
		return nil, err
	}
	return &fileData, nil
}

func ReadBaseItems(filepath string) ([]string, error) {
	data, err := read_items_file(filepath)
	if err != nil {
		return nil, err
	}
	return data.Items.Base, nil
}

func ReadTarotItems(filepath string) ([]string, error) {
	data, err := read_items_file(filepath)
	if err != nil {
		return nil, err
	}
	return data.Items.Tarot, nil
}

func ReadPillsItems(filepath string) ([]string, error) {
	data, err := read_items_file(filepath)
	if err != nil {
		return nil, err
	}
	return data.Items.Pills, nil
}

func search_room(usr *Player) *Room {
	for i := range rooms.Rooms {
		if rooms.Rooms[i].Name == usr.Zone {
			return &rooms.Rooms[i]
		}
	}
	return nil
}

func getToTeleportRoom(room_type string) string {
	for _, room := range rooms.Rooms {
		if room.Type == room_type {
			return room.Name
		}
	}
	return ""
}

func getRandomBossRoom() string {
	var possible_rooms []string

	for _, room := range rooms.Rooms {
		if room.Type == "Boss_room" {
			possible_rooms = append(possible_rooms, room.Name)
		}
	}
	if len(possible_rooms) == 0 {
		return ""
	}
	return possible_rooms[rand.IntN(len(possible_rooms))]
}

func getTargetZone(usr *Player) string {
	playerRoom := search_room(usr)

	if playerRoom != nil {
		hasEnemy := false
		for _, npc := range playerRoom.NPC {
			if npc.Ennemy && npc.HP > 0 {
				hasEnemy = true
				break
			}
		}
		if !hasEnemy {
			return playerRoom.Name
		}
	}

	var safeZones []string
	for _, r := range rooms.Rooms {
		hasEnemy := false
		for _, npc := range r.NPC {
			if npc.Ennemy && npc.HP > 0 {
				hasEnemy = true
				break
			}
		}
		if !hasEnemy {
			safeZones = append(safeZones, r.Name)
		}
	}

	if len(safeZones) > 0 {
		return safeZones[rand.IntN(len(safeZones))]
	}

	return "Spawn"
}

func useActiveItem(usr *Player, item_name string) (string, error) {
	switch item_name {
	case "The Book Of Belial":
		applyBookOfBelial(usr)
	case "Book Of Revelations":
		applyBookOfRevelations(usr)
	case "Deck of Cards":
		return applyDecksOfCards(usr), nil
	case "D6":
		applyD6(usr)
	case "Kamikaze!":
		applyKamikaze(usr)
	case "Mom's Bottle Of Pills":
		applyMomsBottleOfPills(usr)
	case "Mom's Bra":
		applyMomsBra(usr)
	case "The nail":
		applyTheNail(usr)
	case "<3":
		applyHeart(usr)
	case "Blood of martyr":
		applyBloodOfMartyr(usr)
	case "Bobby bomb":
		applyBobbyBomb(usr)
	case "Brimstone":
		applyBrimstone(usr)
	case "Charm Of The Vampire":
		applyCharmOfTheVampire(usr)
	case "Dessert":
		applyDessert(usr)
	case "A Dollar":
		applyADollar(usr)
	case "Growth Hormones":
		applyGrowthHormones(usr)
	case "Halo":
		applyHalo(usr)
	case "Inner":
		applyInner(usr)
	case "Lucky Foot":
		applyLuckyFoot(usr)
	case "Magic mush":
		applyMagicMush(usr)
	case "Max's Head":
		applyMaxsHead(usr)
	case "Money = Power":
		applyMoneyEqualPower(usr)
	case "PHD":
		applyPHD(usr)
	case "A Quarter":
		applyAQuarter(usr)
	case "The Small Rock":
		applyTheSmallRock(usr)
	case "The Fool":
		applyTheFool(usr)
	case "The Emperor":
		applyTheEmperor(usr)
	case "The Lovers":
		applyTheLovers(usr)
	case "Justice":
		applyJustice(usr)
	case "The Hermit":
		applyTheHermit(usr)
	case "XI Strength", "Strength":
		applyStrength(usr)
	case "Death":
		applyDeath(usr)
	case "The Devil":
		applyTheDevil(usr)
	case "The Stars":
		applyTheStars(usr)
	case "The Sun":
		applyTheSun(usr)
	case "?Pill?":
		applyPill(usr)
	default:
		return "", fmt.Errorf("ERR unknown item: %s", item_name)
	}
	return "", nil
}

// BASE ITEMS
func applyBookOfBelial(usr *Player) {
	usr.TempATK += 2
	usr.DMG_Mult += 0.5
}

func applyBookOfRevelations(usr *Player) {
	for i := range online_players.Players {
		if online_players.Players[i].Zone == usr.Zone {
			online_players.Players[i].CURRENT_HP += 20
			if online_players.Players[i].CURRENT_HP > online_players.Players[i].MAX_HP {
				online_players.Players[i].CURRENT_HP = online_players.Players[i].MAX_HP
			}
		}
	}
}

func applyDecksOfCards(usr *Player) string {
	available_tarot, err := ReadTarotItems("data/items.json")
	if err != nil {
		log.Fatal(err)
	}
	randomIndex := rand.IntN(len(available_tarot))
	randomTarot := available_tarot[randomIndex]
	itemDisplayName[randomTarot] = randomTarot
	usr.Inventory = append(usr.Inventory, randomTarot)
	return randomTarot
}

func applyD6(usr *Player) {
	room := search_room(usr)
	if room == nil {
		return
	}
	available_item, err := ReadBaseItems("data/items.json")
	if err != nil {
		log.Fatal(err)
	}
	for i := range room.Items {
		randomIndex := rand.IntN(len(available_item))
		randomItem := available_item[randomIndex]
		itemDisplayName["item."+randomItem] = randomItem
		room.Items[i] = ItemInstance{
			ID:   "item." + randomItem,
			Name: randomItem,
		}
	}
}

func applyKamikaze(usr *Player) {
	usr.CURRENT_HP -= 35
	checkAndRespawnPlayer(usr)
	room := search_room(usr)
	if room == nil {
		return
	}
	for i := range room.NPC {
		if room.NPC[i].Ennemy {
			room.NPC[i].HP -= 35
			removeNpc(room, i)
		}
	}
}

func applyMomsBottleOfPills(usr *Player) {
	available_pills, err := ReadPillsItems("data/items.json")
	if err != nil {
		log.Fatal(err)
	}
	randomIndex := rand.IntN(len(available_pills))
	randomPills := available_pills[randomIndex]
	itemDisplayName[randomPills] = randomPills
	usr.Inventory = append(usr.Inventory, randomPills)
}

func applyMomsBra(usr *Player) {}

func applyTheNail(usr *Player) {
	usr.TempATK += 0.7
}

func applyHeart(usr *Player) {
	usr.MAX_HP += 20
	usr.CURRENT_HP += 20
}

func applyBloodOfMartyr(usr *Player) {
	usr.ATK += 1
}

func applyBobbyBomb(usr *Player) {
	usr.Bomb += 5
}

func applyBrimstone(usr *Player) {
	usr.DMG_Mult += 3
}

func applyCharmOfTheVampire(usr *Player) {
	refill := float32(usr.MAX_HP) * 0.02
	usr.CURRENT_HP += int(refill)
	if usr.CURRENT_HP > usr.MAX_HP {
		usr.CURRENT_HP = usr.MAX_HP
	}
}

func applyDessert(usr *Player) {
	usr.MAX_HP += 20
	usr.CURRENT_HP += 20
}

func applyADollar(usr *Player) {
	usr.Coin = 99
}

func applyGrowthHormones(usr *Player) {
	usr.ATK += 2
}

func applyHalo(usr *Player) {
	usr.ATK += 1
	usr.MAX_HP += 20
}

func applyInner(usr *Player) {
	usr.DMG_Mult += 3
}

func applyLuckyFoot(usr *Player) {
	usr.Luck += 7
}

func applyMagicMush(usr *Player) {
	usr.ATK += 2
	usr.MAX_HP += 20
	usr.DMG_Mult += 1.5
}

func applyMaxsHead(usr *Player) {
	usr.DMG_Mult += 1.5
	usr.ATK += 1
}

func applyMomsKnife(usr *Player) {
	usr.ATK += 15
}

func applyMoneyEqualPower(usr *Player) {
	usr.ATK += float32(usr.Coin)
}

func applyPHD(usr *Player) {
	usr.Inventory = append(usr.Inventory, "?Pill?", "?Pill?")
}

func applyAQuarter(usr *Player) {
	usr.Coin += 25
}

func applyTheSmallRock(usr *Player) {
	usr.ATK += 1
}

// TAROT ITEMS
func applyTheFool(usr *Player) {
	if target := getToTeleportRoom("Spawn_room"); target != "" {
		usr.Zone = target
	}
}
func applyTheEmperor(usr *Player) {
	if target := getRandomBossRoom(); target != "" {
		usr.Zone = target
	}
}

func applyTheLovers(usr *Player) {
	usr.CURRENT_HP += 40
	if usr.CURRENT_HP > usr.MAX_HP {
		usr.CURRENT_HP = usr.MAX_HP
	}
}

func applyJustice(usr *Player) {
	usr.Bomb += 1
	usr.Coin += 1
	usr.Key += 1
	usr.MAX_HP += 20
}

func applyTheHermit(usr *Player) {
	if target := getToTeleportRoom("Shop_room"); target != "" {
		usr.Zone = target
	}
}

func applyStrength(usr *Player) {
	usr.MAX_HP += 20
	usr.TempATK += 3
}

func applyDeath(usr *Player) {
	room := search_room(usr)
	if room == nil {
		return
	}
	for i := range room.NPC {
		if room.NPC[i].Ennemy {
			room.NPC[i].HP = 0
			removeNpc(room, i)
		}
	}
}

func applyTheDevil(usr *Player) {
	usr.ATK += 2
}

func applyTheStars(usr *Player) {
	if target := getToTeleportRoom("Item_room"); target != "" {
		usr.Zone = target
	}
}

func applyTheSun(usr *Player) {
	usr.CURRENT_HP = usr.MAX_HP
	room := search_room(usr)
	if room == nil {
		return
	}
	for i := range room.NPC {
		if room.NPC[i].Ennemy {
			room.NPC[i].HP -= 40
			removeNpc(room, i)
		}
	}
}

// PILLS
func applyPill(usr *Player) {
	effects := []string{
		"hard_penalty_hp", "hard_penalty_atk", "hard_penalty_maxhp",
		"medium_penalty_hp", "medium_penalty_atk", "medium_penalty_maxhp",
		"bonus_hp", "bonus_atk", "bonus_maxhp",
	}
	switch effects[rand.IntN(len(effects))] {
	case "hard_penalty_hp":
		usr.CURRENT_HP = max(0, usr.CURRENT_HP-50)
		checkAndRespawnPlayer(usr)
	case "hard_penalty_atk":
		usr.ATK = max(0, usr.ATK-10)
	case "hard_penalty_maxhp":
		usr.MAX_HP = max(0, usr.MAX_HP-30)
		usr.CURRENT_HP = min(usr.CURRENT_HP, usr.MAX_HP)
		checkAndRespawnPlayer(usr)
	case "medium_penalty_hp":
		usr.CURRENT_HP = max(0, usr.CURRENT_HP-20)
		checkAndRespawnPlayer(usr)
	case "medium_penalty_atk":
		usr.ATK = max(0, usr.ATK-5)
	case "medium_penalty_maxhp":
		usr.MAX_HP = max(0, usr.MAX_HP-15)
		usr.CURRENT_HP = min(usr.CURRENT_HP, usr.MAX_HP)
		checkAndRespawnPlayer(usr)
	case "bonus_hp":
		usr.CURRENT_HP = min(usr.MAX_HP, usr.CURRENT_HP+25)
	case "bonus_atk":
		usr.ATK += 5
	case "bonus_maxhp":
		usr.MAX_HP += 20
		usr.CURRENT_HP += 20
	}
}
