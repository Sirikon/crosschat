package server

// Message from users
type Message struct {
	sender *User
	text   string
}
