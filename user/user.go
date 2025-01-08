package user

import (
	"fmt"
	"net"
)

type UserManager struct {
	name string
}

type UserManagerInterface interface {
	ConnectToServer(string)
	DisconnectFromServer()
	HandleConnection()
}

func (u *UserManager) ConnectToServer(ip string) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Printf("Error while connectiong: %v", err)
	}
	instructBuffer := make([]byte, 100)
	_, err = conn.Read(instructBuffer)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print("instructions: a\n", string(instructBuffer))
	go u.HandleConnection(conn)
}

func (u *UserManager) DisconnectFromServer() {

}

func (u *UserManager) HandleConnection(conn net.Conn) {
}
