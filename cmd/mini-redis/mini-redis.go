package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"githbub.com/miladrahmat/MiniRedis/internal/tools"
)

func main() {

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if (err != nil) {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port :6379")
	
	var database *tools.Database = tools.Database.NewDatabase()

	// Listen for connections
	conn, err := l.Accept()
	if (err != nil) {
		fmt.Println(err)
	}
	
	defer conn.Close()

	for {
		resp := NewResp(conn)

		// Read message from client
		value, err = resp.Read()
		if (err != nil) {
			fmt.Println("Error reading from client: ", err.Error())
			return
		}

		fmt.Println(value)

		// Ignore request, reply with PONG
		conn.Write([]byte("+OK\r\n"))
	}
}