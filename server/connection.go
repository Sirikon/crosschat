package server

type Connection interface {
  Send(msg Message)
  Receive() (string, error)
  Close()
}
