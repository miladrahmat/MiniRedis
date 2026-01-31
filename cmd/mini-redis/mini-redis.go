package main

import (
	"fmt"
	"context"
	"errors"
	"net"
	"io"
	"sync"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"github.com/miladrahmat/MiniRedis/internal/resp"
	"github.com/miladrahmat/MiniRedis/internal/tools"
)

var wg sync.WaitGroup

func main() {
	var ctx, cancel = context.WithCancel(context.Background())

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// In case of panic, making sure the listener closes
	defer l.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigChan
		fmt.Println("Shutting down...")
		signal.Stop(sigChan)
		cancel() // All goroutines stop
		l.Close()
	}()

	fmt.Println("Listening on port :6379")

	// Listen for connections concurrently
	for {
		conn, err := l.Accept()
		if (err != nil) {
			select {
				case <-ctx.Done():
					// Server shut down, expected
				
				default:
					// Unexpected error
					fmt.Println(err)
			}
			break
		}

		wg.Add(1)
		go handleClient(conn, ctx)
	}

	wg.Wait()
}

func handleClient(conn net.Conn, ctx context.Context) {
	defer wg.Done()
	defer conn.Close()

	for {
		// Check if the server is still up
		select {
			case <-ctx.Done():
				return

			default:
				// Execute the loop
		}

		reader := resp.NewResp(conn)

		// Read message from client
		value, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
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