package server

import (
	"fmt"
	"net"
)

type Connections []*net.Conn

type ServerManager struct {
	Listener net.Listener
	ConnList Connections
	//ConnManager ConnectionManager
}

type ServerManagerInterface interface {
	StartServer() error
	CloseServer() error
	AcceptConnection() error
	HandleConnection(conn net.Conn) error
}

func (s *ServerManager) CreateServer() error {
	newListener, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		return err
	}
	fmt.Println("Server succesfully created!")
	s.Listener = newListener
	return nil
}

func (s *ServerManager) CloseServer() error {
	err := s.Listener.Close()
	if err != nil {
		return err
	}
	fmt.Println("Server succesfully created!")
	return nil
}

func (s *ServerManager) AcceptConnection() error {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Printf("Error while connection: %v", err)
			continue
		}
		go s.HandleConnection(&conn)
	}
}

func (s *ServerManager) HandleConnection(conn *net.Conn) error {
	s.ConnList = append(s.ConnList, conn)
	return nil
}
