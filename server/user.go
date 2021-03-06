package server

// User that enters to the server
type User struct {
	conn       Connection
	name       string
	channelIn  chan Message
	channelOut chan Message
}

// SetOutgoingChannel to use
func (u *User) SetOutgoingChannel(channel chan Message) {
	u.channelOut = channel
}

// Handle a connection
func (u *User) Handle(conn Connection) {
	u.conn = conn
	u.channelIn = make(chan Message)

	go u.waitIncomingMessages()
	go u.waitOutgoingMessages()
}

// SendMessage to the user
func (u *User) SendMessage(message Message) {
	u.channelIn <- message
}

func (u *User) waitIncomingMessages() {
	for {
		message := <-u.channelIn
		u.conn.Send(message)
	}
}

func (u *User) waitOutgoingMessages() {
	connectionError := false
	for !connectionError {
		messageText, err := u.conn.Receive() // output message received
		if err != nil {
			u.conn.Close()
			connectionError = true
		} else {
			message := Message{sender: u, text: messageText}
			u.channelOut <- message
		}
	}
}
