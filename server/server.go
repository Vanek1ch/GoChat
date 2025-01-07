package server

import (
	"fmt"
	"net"
	"sync"
)

type Connections []net.Conn

type ServerManager struct {
	Listener  net.Listener
	ConnList  Connections
	ConnMutex sync.Mutex
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
	s.ConnMutex.Lock()
	for _, conn := range s.ConnList {
		conn.Close()
	}
	s.ConnList = nil
	s.ConnMutex.Unlock()
	err := s.Listener.Close()
	if err != nil {
		return err
	}
	fmt.Println("Server succesfully closed!")
	return nil
}

func (s *ServerManager) AcceptConnection() error {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Printf("Error while connection: %v\n", err)
			continue
		}
		go s.HandleConnection(conn)
	}
}

func (s *ServerManager) HandleConnection(conn net.Conn) error {
	defer conn.Close()
	fmt.Printf("New connection from %v", conn.RemoteAddr())

	s.ConnMutex.Lock()
	s.ConnList = append(s.ConnList, conn)
	s.ConnMutex.Unlock()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Connection err form %v: %v\n", conn.RemoteAddr(), err)
			break
		}
		recievedData := string(buffer[:n])
		//change logic
		fmt.Print(recievedData)
	}
	return nil
}
