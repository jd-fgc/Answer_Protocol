package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	inventory.Add(widget.NewButton("USE", func() {
		set_use()
	}))
	inventory.Add(widget.NewButton("INVENTORY", func() {
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
func set_right_container(views []*PlayerView, chatTab *fyne.Container, groupTab *fyne.Container) *fyne.Container {
	set_actions()
	set_inventory()
	set_combat()
	set_log()

	apps := container.NewAppTabs(
		container.NewTabItem("Actions", actions),
		container.NewTabItem("Inventory", inventory),
		container.NewTabItem("Combat", combat),
		container.NewTabItem("Chat", chatTab),
		container.NewTabItem("Group", groupTab),
	)

	quitBtn := widget.NewButton("QUIT", func() {
		netClient.SendCommand("QUIT")
		netClient.Close()
		fyne.CurrentApp().Quit()
	})

	logInner := container.NewVBox(layout.NewSpacer(), log)
	logBox := fixedSize(fyne.NewSize(1, 90), container.NewVScroll(logInner))

	gap := canvas.NewRectangle(color.Transparent)
	gap.SetMinSize(fyne.NewSize(1, 15))

	bottomButtons := container.NewGridWithColumns(1, quitBtn)
	bottomMargin := canvas.NewRectangle(color.Transparent)
	bottomMargin.SetMinSize(fyne.NewSize(1, 20))

	bottomArea := container.NewVBox(gap, logBox, bottomMargin, bottomButtons)

	return container.NewBorder(nil, bottomArea, nil, nil, apps)
}
