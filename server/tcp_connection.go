package server

import (
  "bufio"
  "net"
)

type TCPConnection struct {
	conn       net.Conn
}

func (c *TCPConnection) Send(msg Message) {
  c.conn.Write([]byte(msg.text + "\n"))
}

func (c *TCPConnection) Receive() (string, error) {
  return bufio.NewReader(c.conn).ReadString('\n')
}

func (c *TCPConnection) Close() {
  c.conn.Close()
}
