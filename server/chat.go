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
		return "ERR usage: GROUP <CREATE|INVITE|JOIN|LEAVE> ...\n"
	}

	sub := args[0]
	rest := args[1:]

	switch sub {
	case "CREATE":
		if len(rest) == 0 {
			return "ERR missing group name\n"
		}
		if session.Group != "" {
			return "ERR 402 ALREADY_IN_GROUP\n"
		}
		session.Group = rest[0]
		return "OK group=" + rest[0] + "\n"

	case "INVITE":
		if session.Group == "" {
			return "ERR 401 NOT_IN_GROUP\n"
		}
		if len(rest) == 0 {
			return "ERR missing player name\n"
		}
		for _, s := range sessions {
			if s.Username == rest[0] {
				send(s, "EVT GROUP INVITE "+session.Username+" "+session.Group+"\n")
				return "OK\n"
			}
		}
		return "ERR player not found\n"

	case "JOIN":
		if len(rest) == 0 {
			return "ERR missing group name\n"
		}
		if session.Group != "" {
			return "ERR 402 ALREADY_IN_GROUP\n"
		}
		session.Group = rest[0]
		broadcastGroup(rest[0], "EVT GROUP JOIN "+session.Username+"\n")
		return "OK group=" + rest[0] + "\n"

	case "LEAVE":
		if session.Group == "" {
			return "ERR 401 NOT_IN_GROUP\n"
		}
		old := session.Group
		session.Group = ""
		broadcastGroup(old, "EVT GROUP LEAVE "+session.Username+"\n")
		return "OK\n"

	default:
		return "ERR unknown group subcommand\n"
	}
}
