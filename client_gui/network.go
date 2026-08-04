package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type NetClient struct {
	conn      net.Conn
	responses chan string
	Events    chan string
}

func Connect(ip, port string) (*NetClient, error) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return nil, err
	}

	c := &NetClient{
		conn:      conn,
		responses: make(chan string, 20),
		Events:    make(chan string, 50),
	}

	go func() {
		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				close(c.responses)
				close(c.Events)
				return
			}
			if strings.HasPrefix(line, "EVT ") {
				select {
				case c.Events <- line:
				default:
				}
			} else {
				c.responses <- line
			}
		}
	}()

	return c, nil
}

func (c *NetClient) SendCommand(cmd string) ([]string, error) {
	_, err := c.conn.Write([]byte(cmd + "\n"))
	if err != nil {
		return nil, err
	}

	first, ok := <-c.responses
	if !ok {
		return nil, fmt.Errorf("connexion fermée")
	}
	lines := []string{first}

	for {
		select {
		case extra, ok := <-c.responses:
			if !ok {
				return lines, nil
			}
			lines = append(lines, extra)
		default:
			return lines, nil
		}
	}
}

func (c *NetClient) Close() {
	c.conn.Close()
}

var netClient *NetClient

func doConnect(ip, port, username string) ([]string, error) {
	client, err := Connect(ip, port)
	if err != nil {
		return nil, fmt.Errorf("connexion impossible: %w", err)
	}
	netClient = client
	return netClient.SendCommand("CONNECT " + username)
}
