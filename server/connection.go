package server

import "net"

type Connection struct {
	Conn net.Conn
	Name string
}

type ConnectionManager struct {
	Connections *Connection
}

type ConnectionManagerInterface interface {
	HandleConnection(conn net.Conn)
}

func (c *ConnectionManager) HandleConnection() {

}
