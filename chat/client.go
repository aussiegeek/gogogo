package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

type Client struct {
	Connection net.Conn
	Name       string
}

func (client Client) Send(message string) {
	fmt.Fprintf(client.Connection, message+"\n")
}

func main() {
	flag.Parse()

	connection, err := net.Dial("tcp", "localhost:3337")
	client := Client{Connection: connection, Name: flag.Args()[0]}
	if err != nil {
		log.Fatal("Could not connect to server.")
	}
	client.Send("NAME " + client.Name)
	incoming := bufio.NewScanner(connection)
	go func() {
		for incoming.Scan() {
			text := incoming.Text()
			fmt.Println(text)
		}
	}()

	messages := bufio.NewScanner(os.Stdin)
	for messages.Scan() {
		client.Send("MSG " + messages.Text())
	}
}
