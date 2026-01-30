package main

import (
	"fmt"
	"net"
	"strings"
	"github.com/miladrahmat/MiniRedis/internal/resp"
	"github.com/miladrahmat/MiniRedis/internal/tools"
)

func main() {

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if (err != nil) {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port :6379")
	
	// database := tools.NewDatabase()

	// Listen for connections
	conn, err := l.Accept()
	if (err != nil) {
		fmt.Println(err)
	}
	
	defer conn.Close()

	for {
		reader := resp.NewResp(conn)

		// Read message from client
		value, err := reader.Read()
		if (err != nil) {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err.Error())
			return
		}

		if value.Typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected array length greater than 0")
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)

		handler, ok := tools.Handlers[command]

		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(resp.Value{Typ: "string", Str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}