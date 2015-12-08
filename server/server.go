package server

import "net"
import "fmt"

import "strconv"

// Start the Server service
func Start() {
	fmt.Println("Chat Server 8081")     // listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081") // accept connection on port

	msgIn := make(chan Message)
	msgOut := make(chan Message)
	connections := make(chan net.Conn)

	go waitConnections(ln, connections)
	go waitUsersToHandle(connections, msgIn, msgOut)
	go messageBroadcaster(msgIn, msgOut)

	for {
		// Keep the thread working
	}
}

func messageBroadcaster(msgIn chan Message, msgOut chan Message) {
	for {
		message := <-msgIn
		msgOut <- message
	}
}

func waitConnections(ln net.Listener, connections chan net.Conn) {
	for { // will listen for message to process ending in newline (\n)
		conn, _ := ln.Accept() // run loop forever (or until ctrl-c)
		fmt.Println("New Connection")
		connections <- conn
	}
}

func waitUsersToHandle(connections chan net.Conn, msgIn chan Message, msgOut chan Message) {
	userList := make([]*User, 10)
	userListCount := 0
	for {

		select {
		case conn := <-connections:
			user := &User{}
			user.SetOutgoingChannel(msgIn)
			user.Handle(conn)
			userList[userListCount] = user
			fmt.Println("User stored at " + strconv.Itoa(userListCount))
			userListCount++
		default:
			// ...
		}

		select {
		case msg := <-msgOut:
			for i := 0; i < userListCount; i++ {
				if userList[i] != msg.sender {
					fmt.Println("Sending message to conn #" + strconv.Itoa(i))
					userList[i].SendMessage(msg)
				}
			}
		default:
			// ...
		}

	}
}
