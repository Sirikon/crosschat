package server

import "net"
import "net/http"

import "log"
import "fmt"

import "strconv"

import "golang.org/x/net/websocket"

// Start the Server service
func Start() {
	fmt.Println("Chat Server 8081")     // listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081") // accept connection on port

	msgIn := make(chan Message)
	msgOut := make(chan Message)
	connections := make(chan Connection)

	go waitTCPConnections(ln, connections)
	go waitWSConnections(connections)
	go waitUsersToHandle(connections, msgIn, msgOut)
	go messageBroadcaster(msgIn, msgOut)

	http.Handle("/", http.FileServer(http.Dir("server/webroot")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func messageBroadcaster(msgIn chan Message, msgOut chan Message) {
	for {
		message := <-msgIn
		msgOut <- message
	}
}

func waitTCPConnections(ln net.Listener, connections chan Connection) {
	for { // will listen for message to process ending in newline (\n)
		conn, _ := ln.Accept() // run loop forever (or until ctrl-c)
		fmt.Println("New Connection")
		connections <- &TCPConnection{conn}
	}
}

func waitWSConnections(connections chan Connection) {
	// websocket handler
	onConnected := func(conn *websocket.Conn) {
		connections <- &WSConnection{conn}

		for {}
	}
	http.Handle("/wsusers", websocket.Handler(onConnected))
}

func waitUsersToHandle(connections chan Connection, msgIn chan Message, msgOut chan Message) {
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
