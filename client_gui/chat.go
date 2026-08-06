package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var chatLog *widget.Label

func buildChatTab() *fyne.Container {
	chatLog = widget.NewLabel("")
	chatLog.Wrapping = fyne.TextWrapWord

	scroll := container.NewVScroll(chatLog)

	input := widget.NewEntry()
	input.SetPlaceHolder("Message... (ex: /room hello, /group hey, sinon global)")

	send := func() {
		text := strings.TrimSpace(input.Text)
		if text == "" {
			return
		}
		scope := "GLOBAL"
		if strings.HasPrefix(text, "/room ") {
			scope = "ROOM"
			text = strings.TrimPrefix(text, "/room ")
		} else if strings.HasPrefix(text, "/group ") {
			scope = "GROUP"
			text = strings.TrimPrefix(text, "/group ")
		}

		_, err := netClient.SendCommand("CHAT " + scope + " " + text)
		if err != nil {
			appendChat("Erreur envoi: " + err.Error())
		}
		input.SetText("")
	}

	input.OnSubmitted = func(string) { send() }
	sendBtn := widget.NewButton("Envoyer", send)

	return container.NewBorder(nil, container.NewBorder(nil, nil, nil, sendBtn, input), nil, nil, scroll)
}

func appendChat(line string) {
	if chatLog == nil {
		return
	}
	fyne.Do(func() {
		chatLog.SetText(chatLog.Text + line + "\n")
	})
}

func listenChatEvents() {
	for evt := range netClient.Events {
		trimmed := strings.TrimSpace(evt)
		if strings.HasPrefix(trimmed, "EVT GROUP") {
			appendGroupEvent(trimmed)
		} else {
			appendChat(trimmed)
		}
	}
}
