package main

import (
	"encoding/json"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func doInventory() ([]itemInstance, error) {
	lines, err := netClient.SendCommand("INVENTORY")
	if err != nil || len(lines) == 0 {
		return nil, err
	}

	line := strings.TrimSpace(lines[0])
	line = strings.TrimPrefix(line, "OK ")

	var inv []itemInstance
	if err := json.Unmarshal([]byte(line), &inv); err != nil {
		return nil, err
	}
	return inv, nil
}

func inventory_description() string {
	inv, err := doInventory()
	if err != nil {
		return "Erreur INVENTORY: " + err.Error()
	}
	if len(inv) == 0 {
		return "Inventory: (empty)"
	}
	var names []string
	for _, it := range inv {
		names = append(names, it.Name)
	}
	return "Inventory: " + strings.Join(names, ", ")
}

func set_take() {
	refreshRoomState(false)
	inventory.RemoveAll()

	for _, item := range currentRoomItems {
		item := item
		inventory.Add(widget.NewButton(item.Name, func() {
			lines, err := netClient.SendCommand("TAKE " + item.ID)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du TAKE")
				return
			}
			log_text.SetText(strings.TrimSpace(lines[0]))
			refreshRoomState(false)
			set_take()
		}))
	}
	inventory.Add(widget.NewButton("RETURN", func() {
		set_inventory()
	}))
	inventory.Refresh()
}

func set_drop() {
	refreshRoomState(false)
	inventory.RemoveAll()

	inv, err := doInventory()
	if err != nil {
		log_text.SetText("Erreur INVENTORY: " + err.Error())
	}

	for _, item := range inv {
		item := item
		inventory.Add(widget.NewButton(item.Name, func() {
			lines, err := netClient.SendCommand("DROP " + item.ID)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du DROP")
				return
			}
			log_text.SetText(strings.TrimSpace(lines[0]))
			refreshRoomState(false)
			set_drop()
		}))
	}
	inventory.Add(widget.NewButton("RETURN", func() {
		set_inventory()
	}))
	inventory.Refresh()
}

type roomSummary struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Exits lookExits `json:"exits"`
}

func set_use() {
	refreshRoomState(false)
	inventory.RemoveAll()

	inv, err := doInventory()
	if err != nil {
		log_text.SetText("Erreur INVENTORY: " + err.Error())
	}

	for _, item := range inv {
		item := item
		inventory.Add(widget.NewButton(item.Name, func() {
			lines, err := netClient.SendCommand("USE " + item.ID)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du USE")
				return
			}
			applyUseResponse(strings.TrimSpace(lines[0]))
		}))
	}
	inventory.Add(widget.NewButton("RETURN", func() {
		set_inventory()
	}))
	inventory.Refresh()
}

func applyUseResponse(line string) {
	if !strings.HasPrefix(line, "OK ") {
		log_text.SetText(line)
		return
	}
	jsonPart := strings.TrimPrefix(line, "OK ")

	var summary []roomSummary
	if err := json.Unmarshal([]byte(jsonPart), &summary); err != nil {
		log_text.SetText("OK")
		return
	}

	for _, rs := range summary {
		updateKnownRoom(rs.Name, rs.Exits)
	}
	refreshMiniMap()
	log_text.SetText("Carte révélée !")
	set_inventory()
}
