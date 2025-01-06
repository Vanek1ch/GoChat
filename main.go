package main

import (
	"fmt"
)

func serverStart() {

}
func serverConnect() {}

func main() {
	var userChoose int = -1
	for {
		fmt.Println("Please input your type:\n (0) - Server; (1) - Client")
		_, err := fmt.Scan(&userChoose)
		if err != nil {
			fmt.Print("Wrong insert!")
		} else {
			switch userChoose {
			case 0:
				serverStart()
				return
			case 1:
				serverConnect()
				return
			default:
				fmt.Printf("%v is wrong number!", userChoose)
			}
		}
	}
}
