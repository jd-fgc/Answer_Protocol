package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


// INITIALISE LE CONTAINER DE GAUCHE
func set_left_container(views []*PlayerView) *container.Split {
	return container.NewVSplit(
		container.NewVSplit(views[0].Container,
							views[1].Container),
		container.NewVSplit(views[2].Container,
							views[3].Container),
	)
}

// INITIALISE L'ONGLET ACTION
func set_actions() {
	actions.RemoveAll()
	actions.Add(widget.NewButton("TALK", func() {
		set_talk()
	}))
	actions.Add(widget.NewButton("QUEST", func() {
		set_quests()
	}))
	actions.Refresh()
}

// INITIALISE L'ONGLET INVENTORY
func set_inventory() {
	inventory.RemoveAll()
	inventory.Add(widget.NewButton("TAKE", func() {}))
	inventory.Add(widget.NewButton("DROP", func() {}))
	inventory.Add(widget.NewButton("INVENTORY", func() {}))
	inventory.Refresh()
}

// INITIALISE L'ONGLET COMBAT
func set_combat() {
	combat.RemoveAll()
	combat.Add(widget.NewButton("ATTACK", func() {}))
	combat.Add(widget.NewButton("DEFEND", func() {}))
	combat.Refresh()
}

// INITIALISE LE CONTAINER DE DROITE
func set_right_container(views []*PlayerView) *container.Split {
	set_actions()
	set_inventory()
	set_combat()
	set_log()

	apps := container.NewAppTabs(
		container.NewTabItem("Actions", actions),
		container.NewTabItem("Inventory", inventory),
		container.NewTabItem("Combat", combat),
	)

	return container.NewVSplit(apps, log)
}