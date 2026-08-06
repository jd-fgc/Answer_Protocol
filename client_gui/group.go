package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var groupStatus *widget.Label

func buildGroupTab() *fyne.Container {
	groupStatus = widget.NewLabel("Groupe : aucun")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Nom du groupe")

	createBtn := widget.NewButton("CRÉER", func() {
		name := strings.TrimSpace(nameEntry.Text)
		if name == "" {
			return
		}
		lines, err := netClient.SendCommand("GROUP CREATE " + name)
		if err != nil || len(lines) == 0 {
			groupStatus.SetText("Erreur lors de la création du groupe")
			return
		}
		groupStatus.SetText(strings.TrimSpace(lines[0]))
	})

	joinBtn := widget.NewButton("REJOINDRE", func() {
		name := strings.TrimSpace(nameEntry.Text)
		if name == "" {
			return
		}
		lines, err := netClient.SendCommand("GROUP JOIN " + name)
		if err != nil || len(lines) == 0 {
			groupStatus.SetText("Erreur lors de la connexion au groupe")
			return
		}
		groupStatus.SetText(strings.TrimSpace(lines[0]))
	})

	leaveBtn := widget.NewButton("QUITTER", func() {
		_, err := netClient.SendCommand("GROUP LEAVE")
		if err != nil {
			return
		}
		groupStatus.SetText("Groupe : aucun")
	})

	inviteEntry := widget.NewEntry()
	inviteEntry.SetPlaceHolder("Pseudo à inviter")
	inviteBtn := widget.NewButton("INVITER", func() {
		target := strings.TrimSpace(inviteEntry.Text)
		if target == "" {
			return
		}
		netClient.SendCommand("GROUP INVITE " + target)
	})

	return container.NewVBox(
		groupStatus,
		nameEntry,
		container.NewHBox(createBtn, joinBtn, leaveBtn),
		container.NewHBox(inviteEntry, inviteBtn),
	)
}

func appendGroupEvent(line string) {
	if groupStatus == nil {
		return
	}
	fyne.Do(func() {
		groupStatus.SetText(groupStatus.Text + "\n" + line)
	})
}
