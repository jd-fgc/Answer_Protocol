package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

var online_players Players
var playable_characters Players
var rooms Rooms
var quests Quests

type Message struct {
	From      string
	Text      string
	ReplyChan chan string
}

func hub(incoming chan Message) {
	for {
		msg := <-incoming
		cmd, args := parseCommand(msg.Text)

		playerName := "unknown"
		if s, ok := sessions[msg.From]; ok && s.Username != "" {
			playerName = s.Username
		}

		logInfo("command_received", map[string]interface{}{
			"from": msg.From, "player": playerName, "command": cmd, "args": args,
		})

		if checkAbuse(msg.From) {
			logWarn("command_flooding", map[string]interface{}{"from": msg.From, "player": playerName})
		}
		nolog := false
		var reply string
		switch cmd {
		case "CONNECT":
			reply = handleConnect(msg.From, args)
		case "CHOOSE":
			reply = handleChoose(msg.From, args)
		case "MOVE":
			reply = handleMove(msg.From, args)
		case "LOOK":
			reply = handleLook(msg.From)
		case "TALK":
			reply = handleTalk(msg.From, args)
		case "TAKE":
			reply = handleTake(msg.From, args)
		case "DROP":
			reply = handleDrop(msg.From, args)
		case "INVENTORY":
			reply = handleInventory(msg.From)
		case "USE":
			reply = handleUse(msg.From, args)
		case "WHO":
			reply = handleWho(msg.From)
		case "CHAT":
			reply = handleChat(msg.From, args)
		case "GROUP":
			reply = handleGroup(msg.From, args)
		case "QUIT":
			reply = handleQuit(msg.From)
		case "ATTACK":
			reply = handleAttack(msg.From, args)
		case "DEFEND":
			reply = handleDefend(msg.From)
		case "STATUS":
			reply = handleStatus(msg.From)
		case "QUEST":
			reply = handleQuest(msg.From, args)
		case "QUESTS":
			reply = handleQuests(msg.From)
		case "COMPLETE_QUEST":
			reply = handleCompleteQuest(msg.From, args)
		case "REFRESH_ROOMS":
			reply = REFRESH_ROOMS()
			nolog = true
		default:
			reply = "ERR unknown command\n"
		}

		if nolog == false {
			logInfo("response_sent", map[string]interface{}{
				"from": msg.From, "player": playerName, "response": strings.TrimSpace(reply),
			})
		}
		

		msg.ReplyChan <- reply
	}
}

func parseCommand(line string) (string, []string) {
	parts := strings.Fields(strings.TrimSpace(line))
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}

func getConnectedPlayer(from string) (*Player, error) {
	session, ok := sessions[from]
	if !ok || session.Player == nil {
		return nil, errors.New("not connected")
	}
	return session.Player, nil
}

func handleConnection(conn net.Conn, incoming chan Message) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	from := conn.RemoteAddr().String()

	conn.Write([]byte("OK hello proto=1\n"))

	stop := make(chan struct{})
	go func() {
		for {
			session, ok := sessions[from]
			if !ok || session.Events == nil {
				select {
				case <-stop:
					return
				case <-time.After(50 * time.Millisecond):
					continue
				}
			}
			select {
			case evt := <-session.Events:
				conn.Write([]byte(evt))
			case <-stop:
				return
			}
		}
	}()

	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				logError("read_failed", map[string]interface{}{"from": from, "error": err.Error()})
			}
			logInfo("client_disconnected", map[string]interface{}{"from": from})

			replyChan := make(chan string, 1)
			incoming <- Message{From: from, Text: "QUIT\n", ReplyChan: replyChan}
			close(stop)
			return
		}

		replyChan := make(chan string)
		incoming <- Message{From: from, Text: string(bytes), ReplyChan: replyChan}
		response := <-replyChan
		conn.Write([]byte(response))

		cmd, _ := parseCommand(string(bytes))
		if cmd == "QUIT" {
			close(stop)
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		os.Exit(1)
	}

	if err := initLogger(); err != nil {
		fmt.Println("warning: logger init failed, falling back to stdout:", err)
	}

	var err error

	port := fmt.Sprintf(":%s", os.Args[1])
	rooms, err = load_map("data/world_map.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	playable_characters, err = read_character_file("data/Characters.json", rooms.Rooms[0].Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	quests, err = load_quests("data/quests.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		logError("listener_failed", map[string]interface{}{"error": err.Error()})
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n", listener.Addr())

	incoming := make(chan Message)
	go hub(incoming)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logError("accept_failed", map[string]interface{}{"error": err.Error()})
			continue
		}

		ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
		if checkConnectionAbuse(ip) {
			logWarn("rapid_connection_pattern", map[string]interface{}{"ip": ip})
		}
		logInfo("client_connected", map[string]interface{}{"remote_addr": conn.RemoteAddr().String()})

		go handleConnection(conn, incoming)
	}
}
