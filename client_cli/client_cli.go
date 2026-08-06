package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var ip string
var port string

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run client_cli.go <address> <port>")
		os.Exit(1)
	}

	ip = os.Args[1]
	port = ":" + os.Args[2]
	addr := ip + port

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Connection failed, err: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	greeting, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR failed to read greeting, err: ", err)
		os.Exit(1)
	}
	fmt.Print(greeting)

	responses := make(chan string, 20)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				close(responses)
				return
			}
			if strings.HasPrefix(line, "EVT ") {
				fmt.Print("\n" + line)
			} else {
				responses <- line
			}
		}
	}()

	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>>")
		input, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR failed to read input, err: ", err)
			break
		}
		conn.Write([]byte(input))

		trimmed := strings.TrimSpace(input)
		if trimmed == "exit" || trimmed == "QUIT" {
			break
		}

		response, ok := <-responses
		if !ok {
			fmt.Println("Connexion fermée par le serveur")
			break
		}
		fmt.Print(response)

	drain:
		for {
			select {
			case extra, ok := <-responses:
				if !ok {
					break drain
				}
				fmt.Print(extra)
			default:
				break drain
			}
		}
	}
}

func listenEvents(connReader *bufio.Reader) {
	for {
		line, err := connReader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.HasPrefix(line, "EVT ") {
			fmt.Print("\n" + line)
			fmt.Print(">>>")
		}
	}
}
