package user

import (
	"fmt"
	"net"
)

// Interface implements user's availible actions
type UserInterface interface {
	Connect(net.Conn) error
	Disconnect(net.Conn) error
	MakeMessage(net.Conn) error
}

// User struct (add information..)
type User struct {
	name string
}

func (u *User) String() string {
	return fmt.Sprintf("This is user with name %v", u.name)
}

func (u *User) Connect(conn net.Conn) error {
	return nil
}

func (u *User) Disconnect(conn net.Conn) error {

	return nil
}

func (u *User) MakeMessage(conn net.Conn) error {
	return nil
}

// Very simple name validation
func CreateUser(name string) (User, error) {
	if len(name) > 10 && len(name) < 3 {
		return User{}, nil
	} else {
		newUser := &User{name: name}
		return *newUser, nil
	}
}
