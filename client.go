package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server: ", err)
			return
		}
		fmt.Println(msg)
	}
}

func write(conn net.Conn) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter your message: ")
		msg, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error writing to server: ", err)
		}
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8031", "IP:port string to connect to")
	flag.Parse()

	conn, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}

	go read(conn)
	write(conn)

}
