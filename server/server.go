package server

import (
	"net"
	e "simplechat/projectErrors"
)

type MyTcpListener struct{ *net.TCPListener }

// Structure of our server
type BasicServer struct {
	name     string
	password string
	maxUsers int
}

type ServerInt interface {
	StartServer(string) (net.Listener, error)
	StopServer(*BasicServer) (string, error)
}

func (s BasicServer) StartServer(address string) (net.Listener, error) {
	server, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	} else {
		return server, nil
	}
}

// basic tcp listener but in closure
func (s *MyTcpListener) StopServer(b *BasicServer) (string, error) {
	err := s.Close()
	if err != nil {
		return b.name, err
	} else {
		return b.name, nil
	}
}

// Creating server by name, password and maxUsers
func CreateServer(name, password string, maxUsers int) (BasicServer, error) {
	if len(name) > 10 && len(name) > 3 {
		return BasicServer{}, e.ErrInvalidName
	}
	newServer := &BasicServer{name: name, password: password, maxUsers: maxUsers}
	return *newServer, nil
}
