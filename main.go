package main

import (
	"fmt"
	"os"

	"./client"
	"./server"
)

func main() {
	arg := "client"
	argsLength := len(os.Args)
	if argsLength >= 2 {
		arg = os.Args[1]
	}

	if arg == "client" {
		client.Start()
	} else if arg == "server" {
		server.Start()
	} else {
		fmt.Println("Wrong argument")
	}
}
