package main

import (
	"encoding/json"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func doInventory() ([]string, error) {
	lines, err := netClient.SendCommand("INVENTORY")
	if err != nil || len(lines) == 0 {
		return nil, err
	}
	var inv []string
	if err := json.Unmarshal([]byte(lines[0]), &inv); err != nil {
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
	return "Inventory: " + strings.Join(inv, ", ")
}

func set_take() {
	inventory.RemoveAll()

	for _, item := range currentRoomItems {
		item := item
		inventory.Add(widget.NewButton(item, func() {
			lines, err := netClient.SendCommand("TAKE " + item)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du TAKE")
				return
			}
			log_text.SetText(strings.TrimSpace(lines[0]))
			refreshRoomState(false)
			set_inventory()
		}))
	}
	inventory.Add(widget.NewButton("RETURN", func() {
		set_inventory()
	}))
	inventory.Refresh()
}

func set_drop() {
	inventory.RemoveAll()

	inv, err := doInventory()
	if err != nil {
		log_text.SetText("Erreur INVENTORY: " + err.Error())
	}

	for _, item := range inv {
		item := item
		inventory.Add(widget.NewButton(item, func() {
			lines, err := netClient.SendCommand("DROP " + item)
			if err != nil || len(lines) == 0 {
				log_text.SetText("Erreur réseau lors du DROP")
				return
			}
			log_text.SetText(strings.TrimSpace(lines[0]))
			refreshRoomState(false)
			set_inventory()
		}))
	}
	inventory.Add(widget.NewButton("RETURN", func() {
		set_inventory()
	}))
	inventory.Refresh()
}
