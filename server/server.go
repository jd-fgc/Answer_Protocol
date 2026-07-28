package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

type Message struct {
	From      string
	Text      string
	ReplyChan chan string
}

func hub(incoming chan Message) {
	for {
		msg := <-incoming
		fmt.Printf("request: %s", msg.Text)
		reply := msg.Text
		msg.ReplyChan <- reply
	}
}

func handleConnection(conn net.Conn, incoming chan Message) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}

		replyChan := make(chan string)
		incoming <- Message{From: conn.RemoteAddr().String(), Text: string(bytes), ReplyChan: replyChan}
		response := <-replyChan

		conn.Write([]byte(response))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		os.Exit(1)
	}

	port := fmt.Sprintf(":%s", os.Args[1])

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to create listenet, err:", err)
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
