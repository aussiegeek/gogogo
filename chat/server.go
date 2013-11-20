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
			server.Broadcast(&user, user.Name+" has connected.")
		case "MSG":
			server.Broadcast(&user, strings.Join(args, " "))
			// else:
			// handle unknown messages
		}
	}
}

func (server *Server) Broadcast(sender *User, message string) {
	for _, user := range server.Users {
		if user != sender {
			user.Send(sender.Name + ": " + message + "\n")
		}
	}
}

type User struct {
	Incoming *bufio.Scanner
	Outgoing *bufio.Writer
	Name     string
}

func (user *User) Send(message string) {
	user.Outgoing.WriteString(message)
	user.Outgoing.Flush()
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
