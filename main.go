package main

import (
	"fmt"
	"os"

	"./client"
	"./server"
)

func main() {
	mode := "client"
	address := "127.0.0.1:8081"

	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}

	if len(os.Args) >= 3 {
		address = os.Args[2]
	}

	if mode == "client" {
		client.Start(address)
	} else if mode == "server" {
		server.Start()
	} else {
		fmt.Println("Wrong argument")
	}
}
