package main

import (
	"fmt"
	serv "simplechat/server"
	usr "simplechat/user"
)

func main() {
	var userChoose int = -1
	for {
		fmt.Println("Please input your type:\n (0) - Server; (1) - Client")
		_, err := fmt.Scan(&userChoose)
		if err != nil {
			fmt.Println("Wrong insert!")
			continue
		} else {
			switch userChoose {
			case 0:
				// Need to handle the conn.
				servManager := serv.ServerManager{}
				err := servManager.CreateServer()
				if err != nil {
					fmt.Print(err)
				}
			case 1:
				user := usr.UserManager{}
				user.ConnectToServer("127.0.0.1:12345")
			default:
				fmt.Println(userChoose, " is wrong number!")
			}
		}

	}
}
