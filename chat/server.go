package main

import (
	"bufio"
	"net"
	"strings"
)

type Server struct {
	Users []*User
}

func (server *Server) newConnection(conn net.Conn) {
	incoming := bufio.NewScanner(conn)
	outgoing := bufio.NewWriter(conn)
	user := User{Incoming: incoming, Outgoing: outgoing}
	server.Users = append(server.Users, &user)
	for user.Incoming.Scan() {
		input := user.Incoming.Text()
		parts := strings.Split(input, " ")
		command, args := parts[0], parts[1:]
		switch command {
		case "NAME":
			user.Name = args[0]
			server.Broadcast("SERVER", user.Name+" has connected.")
		case "MSG":
			server.Broadcast(user.Name, strings.Join(args, " "))
			// else:
			// handle unknown messages
		}
	}
}

func (server *Server) Broadcast(sender string, message string) {
	for _, user := range server.Users {
		if user.Name != sender {
			user.Outgoing.WriteString(sender + ": " + message + "\n")
			user.Outgoing.Flush()
		}
	}
}

type User struct {
	Incoming *bufio.Scanner
	Outgoing *bufio.Writer
	Name     string
}

func main() {
	ln, err := net.Listen("tcp", ":3337")
	server := new(Server)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go server.newConnection(conn)
	}
}
