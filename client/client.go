package client

import "net"
import "fmt"
import "bufio"
import "os"

// Start the Client service
func Start(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Couldn't connect to " + address);
	} else {
		fmt.Println("--- Connected to " + address + " ---")
		go waitForMessages(conn)
		go waitToSendMessages(conn)

		for {
		}
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
