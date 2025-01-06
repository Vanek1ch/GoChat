package server

import (
	"net"
)

type MyTcpListener struct{ *net.TCPListener }

// Structure of our server
type BasicServer struct {
	name string
}

type Server interface {
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

func (s *MyTcpListener) StopServer(b *BasicServer) (string, error) {
	err := s.Close()
	if err != nil {
		return b.name, err
	} else {
		return b.name, nil
	}
}
