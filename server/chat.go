package main

import "strings"

func handleChat(from string, args []string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) < 2 {
		return "ERR usage: CHAT <GLOBAL|ROOM|GROUP> <message>\n"
	}

	scope := args[0]
	message := strings.Join(args[1:], " ")
	event := "EVT " + scope + " CHAT " + session.Username + " " + message + "\n"

	switch scope {
	case "GLOBAL":
		broadcastGlobal(event)
	case "ROOM":
		broadcastRoom(session.Player.Zone, event)
	case "GROUP":
		if session.Group == "" {
			return "ERR not in a group\n"
		}
		broadcastGroup(session.Group, event)
	default:
		return "ERR unknown scope\n"
	}

	return "OK\n"
}

func handleGroup(from string, args []string) string {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return "ERR not connected\n"
	}
	if len(args) == 0 {
		if session.Group == "" {
			return "ERR not in a group\n"
		}
		return "OK group=" + session.Group + "\n"
	}

	switch args[0] {
	case "LEAVE":
		if session.Group == "" {
			return "ERR not in a group\n"
		}
		old := session.Group
		session.Group = ""
		broadcastGroup(old, "EVT GROUP LEAVE "+session.Username+"\n")
		return "OK left\n"
	default:
		groupName := args[0]
		if session.Group == groupName {
			return "ERR already in this group\n"
		}
		if session.Group != "" {
			broadcastGroup(session.Group, "EVT GROUP LEAVE "+session.Username+"\n")
		}
		session.Group = groupName
		broadcastGroup(groupName, "EVT GROUP JOIN "+session.Username+"\n")
		return "OK group=" + groupName + "\n"
	}
}
