package main

import (
	"fmt"
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
				return
			case 1:
				return
			default:
				fmt.Println(userChoose, " is wrong number!")
			}
		}

	}
}
