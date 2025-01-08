package server

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

// net.Conn / IP or Username.
type Connections map[net.Conn]any

type ServerManager struct {
	Listener  net.Listener
	ConnList  Connections
	ConnMutex sync.Mutex
	ChManager *ChannelManager
	CommaList []string
	//ConnManager ConnectionManager
}

type ServerManagerInterface interface {
	StartServer() error
	CloseServer() error
	AcceptConnection() error
	SendInstructions(conn net.Conn) error
	HandleConnection(conn Connections) error
}

func (s *ServerManager) CreateServer() error {
	newListener, err := net.Listen("tcp", ":12345")
	if err != nil {
		return err
	}
	fmt.Println("Server succesfully created!")
	s.Listener = newListener
	s.AcceptConnection()
	return nil
}

func (s *ServerManager) CloseServer() error {
	s.ConnMutex.Lock()
	for conn := range s.ConnList {
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

// Sending simple instructions to user.
func (s *ServerManager) SendInstructions(conn net.Conn) error {
	instructions := []byte("Welcome to GoChat ver 0.1;\n To start working type /help.")
	_, err := conn.Write(instructions)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerManager) HandleConnection(conn net.Conn) error {
	defer func() {
		conn.Close()
		fmt.Printf("Connection closed: %v\n", conn.RemoteAddr())
		s.ConnMutex.Lock()
		delete(s.ConnList, conn)
		s.ConnMutex.Unlock()
	}()

	fmt.Printf("New connection from %v", conn.RemoteAddr())

	s.ConnMutex.Lock()
	s.ConnList[conn] = conn.RemoteAddr().String()
	s.ConnMutex.Unlock()

	err := s.SendInstructions(conn)
	if err != nil {
		return err
	}

	buffer := make([]byte, 256)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Connection err form %v: %v\n", conn.RemoteAddr(), err)
			break
		}

		recievedData := strings.TrimSpace(string(buffer[:n]))
		if len(recievedData) == 0 {
			continue
		}

		switch {
		case strings.HasPrefix(recievedData, "/help"):
			// Need to make list of commands to use with %v.
			conn.Write([]byte("Availible commands(/) are name, showch"))

		// Case to change username.
		case strings.HasPrefix(recievedData, "/name"):
			newName := strings.TrimSpace(strings.TrimPrefix(recievedData, "/name"))
			if newName == "" {
				conn.Write([]byte("ERROR Invalid username.\n"))
				continue
			}
			s.ConnMutex.Lock()
			s.ConnList[conn] = newName
			s.ConnMutex.Unlock()
			conn.Write([]byte(fmt.Sprintf("SUCCESS username changed to: %v", newName)))

		// Case to show channels.
		case strings.HasPrefix(recievedData, "/showch"):
			s.ConnMutex.Lock()
			list := ""
			for channelName := range s.ChManager.List {
				list += (channelName + " ")
			}
			s.ConnMutex.Unlock()
			conn.Write([]byte("Available channels: " + list + "\n"))

		default:
			conn.Write([]byte("Unrecognized command.\n"))
		}
	}
	return nil
}
