package main

import (
	"fmt"
)
func main() {
	var conferenceName string = "Go Conference"
	// conferenceName := "Go Conference"

	const conferenceTicket int = 50
	// not accept negative value
	var remainingTicket uint = 50

	fmt.Printf("Conference name is type %T, Remaining ticket is %T\n", conferenceName, remainingTicket)
	fmt.Printf("Welcom to our %v booking application...\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your ticket here to attend")

	var userName string
	var userTicket int
	userName = "yobicod"
	userTicket = 2

	fmt.Printf("User %v book %v user ticket\n", userName, userTicket)

}
