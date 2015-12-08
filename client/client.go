package client

import "net"
import "fmt"
import "bufio"
import "os"

// Start the Client service
func Start() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	go waitForMessages(conn)
	go waitToSendMessages(conn)

	for {
	}
}

func waitForMessages(conn net.Conn) {
	for { // read in input from stdin
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
}

func waitToSendMessages(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n') // send to socket
		fmt.Fprintf(conn, text+"\n")       // listen for reply
	}
}
