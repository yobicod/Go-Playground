package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTicket int = 50

var remainingTicket uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

var wg = sync.WaitGroup{}

func main() {

	// call greet user msg
	greetUsers()

	// for remainingTicket > 0 && len(bookings) < 50 {
	// get user input
	firstName, lastName, email, userTicket := getUserInput()

	// validate user input
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

	if isValidName && isValidEmail && isValidTicketNumber {
		// call book ticket
		bookTicket(userTicket, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)
		// call getFirstName to print logs
		firstNames := getFirstName()
		fmt.Printf("These first name of all bookings in our app: %v\n", firstNames)

		if remainingTicket == 0 {
			// end program
			fmt.Printf("Our %v is booked out. Comback later!", conferenceName)
			// break
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
		// continue
	}

	wg.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to our %v\n", conferenceName)
	fmt.Printf("Conference name is type %T\nRemaining ticket is %T\n", conferenceName, remainingTicket)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your ticket here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// func validateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

// 	return isValidName, isValidEmail, isValidTicketNumber
// }

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

	// create a map for a user (obj)
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTicket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation at email: %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("Remains ticket available now: %v %v\n", remainingTicket, conferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTicket, firstName, lastName)
	fmt.Println("++++++++++")
	fmt.Printf("Sending tickets: %v \nto email address %v\n", ticket, email)
	fmt.Println("++++++++++")

	wg.Done()
}
