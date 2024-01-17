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

	var firstName string
	var lastName string
	var email string
	
	var userTicket uint

	//input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	
	fmt.Println("Enter tickets number: ")
	fmt.Scan(&userTicket)

	remainingTicket -= userTicket
	
	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation at email: %v", firstName, lastName, userTicket,email)
	fmt.Printf("Remains ticket available now: %v %v\n", remainingTicket, conferenceName)
}
