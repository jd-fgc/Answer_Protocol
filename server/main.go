package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var online_players Players
var playable_characters Players
var rooms Rooms

type Message struct {
	From      string
	Text      string
	ReplyChan chan string
}

func hub(incoming chan Message) {
	for {
		msg := <-incoming
		cmd, args := parseCommand(msg.Text)

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
		case "QUIT":
			reply = handleQuit(msg.From)
		case "WHO":
			reply = WHO(msg.From)
		// case "CHAT":
		// 	reply = CHAT(msg.From)
		default:
			reply = "ERR unknown command\n"
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

	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			replyChan := make(chan string, 1)
			incoming <- Message{From: from, Text: "QUIT\n", ReplyChan: replyChan}
			return
		}

		replyChan := make(chan string)
		incoming <- Message{From: from, Text: string(bytes), ReplyChan: replyChan}
		response := <-replyChan

		conn.Write([]byte(response))

		cmd, _ := parseCommand(string(bytes))
		if cmd == "QUIT" {
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		os.Exit(1)
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

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n", listener.Addr())

	incoming := make(chan Message)
	go hub(incoming)

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Failed to accep connection, err:", err)
			continue
		}

		go handleConnection(conn, incoming)
	}
}
