package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
	fmt.Println("my error!!!!")
	panic(err)
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	conn, err := ln.Accept()
	if err != nil {
		handleError(err)
	}
	conns <- conn

}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	reader := bufio.NewReader(client)
	for {
		msg, err := reader.ReadString('\n')
		newMsg := Message{
			sender:  clientid,
			message: msg,
		}
		msgs <- newMsg
		if err != nil {
			handleError(err)
		}
	}

}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8031", "port to listen on")
	flag.Parse()
	ln, err := net.Listen("tcp", *portPtr)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
	clientid := 0
	if err != nil {
		handleError(err)
	}
	for {
		select {
		case conn := <-conns:
			clients[clientid] = conn
			handleClient(conn, clientid, msgs)
			clientid += 1

		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for id, conn := range clients {
				if msg.sender != id {
					fmt.Fprintln(conn, msg.message)
				}
			}

		}
	}

}
