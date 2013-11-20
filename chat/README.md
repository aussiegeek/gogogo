# Chat

This project provides a crude chat implementation using a server program and a client program. To run:

    go run server.go
    go run client.go Client1
    go run client.go Client2
    <type + hit enter>

You'll see messages typed on Client 2 appearing on Client 1 and vice versa. All messages will appear on the clients they weren't sent from. 