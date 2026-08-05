package main

func send(s *Session, event string) {
	if s.Events == nil {
		return
	}
	select {
	case s.Events <- event:
	default:
	}
}

func broadcastGlobal(event string) {
	for _, s := range sessions {
		send(s, event)
	}
}

func broadcastRoom(zone string, event string) {
	for _, s := range sessions {
		if s.Player != nil && s.Player.Zone == zone {
			send(s, event)
		}
	}
}

func broadcastGroup(group string, event string) {
	for _, s := range sessions {
		if s.Group == group {
			send(s, event)
		}
	}
}
