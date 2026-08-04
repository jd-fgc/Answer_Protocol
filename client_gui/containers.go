package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var leftColumnGlobal *fyne.Container

func fixedSize(size fyne.Size, content fyne.CanvasObject) fyne.CanvasObject {
	sizer := canvas.NewRectangle(color.Transparent)
	sizer.SetMinSize(size)
	return container.NewStack(sizer, content)
}

func set_left_container() *fyne.Container {
	return container.NewVBox()
}

// INITIALISE L'ONGLET ACTION
func set_actions() {
	actions.RemoveAll()
	actions.Add(widget.NewButton("MOVE", func() {
		set_move()
	}))
	actions.Add(widget.NewButton("TALK", func() {
		set_talk()
	}))
	actions.Add(widget.NewButton("LOOK", func() {
		refreshRoomState(true)
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
	inventory.Add(widget.NewButton("TAKE", func() {
		set_take()
	}))
	inventory.Add(widget.NewButton("DROP", func() {
		set_drop()
	}))
	inventory.Add(widget.NewButton("INVENTORY", func() {
		answer.cmd = "INVENTORY"
		log_text.SetText(inventory_description())
	}))
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

	chatBtn := widget.NewButton("CHAT", func() {
		log_text.SetText("Chat : pas encore branché")
	})

	thirdBtn := widget.NewButton("?", func() {})
	thirdBtn.Disable()

	quitBtn := widget.NewButton("QUIT", func() {
		fyne.CurrentApp().Quit()
	})

	bottomButtons := container.NewGridWithColumns(3, chatBtn, thirdBtn, quitBtn)
	bottomArea := container.NewVBox(log, bottomButtons)

	return container.NewBorder(nil, bottomArea, nil, nil, apps)
}
