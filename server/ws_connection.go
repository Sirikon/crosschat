package server

import (
  "fmt"
  "golang.org/x/net/websocket"
)

type WSConnection struct {
	conn       *websocket.Conn
}

type WSIncomingMessage struct {
  Body string `json:"body"`
}

type WSOutgoingMessage struct {
  User string `json:"user"`
  Body string `json:"body"`
}

func (c *WSConnection) Send(msg Message) {
  wsmsg := WSOutgoingMessage{User: msg.sender.name, Body: msg.text}
  websocket.JSON.Send(c.conn, wsmsg)
}

func (c *WSConnection) Receive() (string, error) {
  var msg WSIncomingMessage
	err := websocket.JSON.Receive(c.conn, &msg)
  return msg.Body, err
}

func (c *WSConnection) Close() {
  c.conn.Close()
}
