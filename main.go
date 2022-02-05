package main

import (
	validator "BookingApp/validators"
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Confrence"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		if !validator.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets) {
			continue
		}
		bookTickets(userTickets, firstName, lastName, email)
		//fmt.Printf("User Name %v \nNo of tickets booked %v \n", userName, userTickets)
		firstNames := getFirstNames()
		fmt.Printf("First Names of bookings %v\n", firstNames)
		fmt.Printf("Current Bookings %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Println("Sold out")
			break
		}
	}
}

func greetUsers() {

	fmt.Printf("Welcome to %v booking App \n", conferenceName)
	fmt.Println("We have total ", conferenceTickets, "tickets and", remainingTickets, "are avialable")
	fmt.Println("You can book tickets here")
}
func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)[0]
		firstNames = append(firstNames, names)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name :")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name :")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address :")
	fmt.Scan(&email)

	fmt.Print("Enter no of tickets you want :")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets %v\n", remainingTickets)
}
