package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTicket int = 50

var remainingTicket uint = 50
var bookings = []string{}

func main() {

	// call greet user msg
	greetUsers()

	for remainingTicket > 0 && len(bookings) < 50 {
		// get user input
		firstName, lastName, email, userTicket := getUserInput()

		// validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remainingTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			// call book ticket
			bookTicket(userTicket, firstName, lastName, email)

			// call getFirstName to print logs
			firstNames := getFirstName(bookings)
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

func greetUsers() {
	fmt.Printf("Welcome to our %v\n", conferenceName)
	fmt.Printf("Conference name is type %T\nRemaining ticket is %T\n", conferenceName, remainingTicket)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your ticket here to attend")
}

func getFirstName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTicket -= userTicket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation at email: %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("Remains ticket available now: %v %v\n", remainingTicket, conferenceName)
}
