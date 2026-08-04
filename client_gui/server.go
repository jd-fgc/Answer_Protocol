package main

import (
	"bufio"
	"fmt"
	"net"
)

type NetClient struct {
	conn   net.Conn
	reader *bufio.Reader
}

func Connect(ip, port string) (*NetClient, error) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return nil, err
	}
	return &NetClient{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}, nil
}

// send command = maybe plusieurs ligne
func (c *NetClient) SendCommand(cmd string) ([]string, error) {
	_, err := c.conn.Write([]byte(cmd + "\n"))
	if err != nil {
		return nil, err
	}

	var lines []string
	line, err := c.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	lines = append(lines, line)

	for c.reader.Buffered() > 0 {
		extra, err := c.reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, extra)
	}

	return lines, nil
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
