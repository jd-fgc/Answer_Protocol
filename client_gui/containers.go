package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Force un CanvasObject à occuper exactement `size`, peu importe son contenu
func fixedSize(size fyne.Size, content fyne.CanvasObject) fyne.CanvasObject {
	sizer := canvas.NewRectangle(color.Transparent)
	sizer.SetMinSize(size)
	return container.NewStack(sizer, content)
}

func set_left_container(views []*PlayerView) *fyne.Container {
	return container.NewGridWithRows(4,
		views[0].Container,
		views[1].Container,
		views[2].Container,
		views[3].Container,
	)
}

// INITIALISE L'ONGLET ACTION
func set_actions() {
	actions.RemoveAll()
	actions.Add(widget.NewButton("LOOK", func() {
		answer.cmd = "LOOK"
		log_text.SetText(look_room_description())
	}))
	actions.Add(widget.NewButton("MOVE", func() {
		set_move()
	}))
	actions.Add(widget.NewButton("TALK", func() {
		set_talk()
	}))
	actions.Add(widget.NewButton("TAKE", func() {
		set_take()
	}))
	actions.Add(widget.NewButton("DROP", func() {
		set_drop()
	}))
	actions.Add(widget.NewButton("INVENTORY", func() {
		answer.cmd = "INVENTORY"
		log_text.SetText(inventory_description())
	}))
	actions.Add(widget.NewButton("WHO", func() {
		answer.cmd = "WHO"
		log_text.SetText(who_description())
	}))
	actions.Add(widget.NewButton("QUEST", func() {
		set_quests()
	}))
	actions.Add(widget.NewButton("QUIT", func() {
		answer.cmd = "QUIT"
		log_text.SetText("OK bye")
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
func set_right_container(views []*PlayerView) *fyne.Container {
	set_actions()
	set_inventory()
	set_combat()
	set_log()

	apps := container.NewAppTabs(
		container.NewTabItem("Actions", actions),
		container.NewTabItem("Inventory", inventory),
		container.NewTabItem("Combat", combat),
	)

	return container.NewBorder(nil, log, nil, nil, apps)
}
