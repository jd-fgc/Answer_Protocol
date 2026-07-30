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
		fmt.Println("Usage: go run client.go <address> <port>")
		os.Exit(1)
	}

	ip = fmt.Sprintf("%s", os.Args[1])
	port = fmt.Sprintf(":%s", os.Args[2])
	addr := ip + port

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Connection failed, err: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	stdinReader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	for {
		fmt.Print(">>>")
		input, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR failed to read input, err: ", err)
			break
		}
		conn.Write([]byte(input))
		input = strings.TrimSpace(input)
		if input == "EXIT" || input == "QUIT" {
			os.Exit(1)
		}

		response, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR failed to read response, err: ", err)
		}
		fmt.Print(response)

		for connReader.Buffered() > 0 {
			extra, err := connReader.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print(extra)
		}
	}
}
