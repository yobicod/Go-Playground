package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	// conferenceName := "Go Conference"

	const conferenceTicket int = 50
	// not accept negative value
	var remainingTicket uint = 50
	var bookings = []string{}

	fmt.Printf("Conference name is type %T, Remaining ticket is %T\n", conferenceName, remainingTicket)
	fmt.Printf("Welcom to our %v booking application...\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your ticket here to attend")

	for remainingTicket > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTicket uint
		// define arr size

		//input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter tickets number: ")
		fmt.Scan(&userTicket)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTicket -= userTicket
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation at email: %v", firstName, lastName, userTicket, email)
			fmt.Printf("Remains ticket available now: %v %v\n", remainingTicket, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var name = strings.Fields(booking)
				firstNames = append(firstNames, name[0])
			}
			fmt.Printf("These first name of all bookings in our app: %v\n", firstNames)

			if remainingTicket == 0 {
				// end program
				fmt.Printf("Our %v is booked out. Comback later!", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Your firstname or lastname is too short\n")
			}

			if !isValidEmail {
				fmt.Printf("Your email not contain @ \n")
			}

			if !isValidTicketNumber {
				fmt.Printf("Wrong ticker number\n")
			}
			continue
		}
	}
}
