package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Confrence"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking App \n", conferenceName)
	fmt.Println("We have total ", conferenceTickets, "tickets and", remainingTickets, "are avialable")
	fmt.Println("You can book tickets here")

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		if !isValidName {
			fmt.Println("Invalid Name")
			continue
		}
		if !isValidEmail {
			fmt.Println("Invalid Email")
			continue
		}
		if isValidTicket {
			fmt.Println("Invalid Ticket")
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("Remaining tickets %v\n", remainingTickets)
		//fmt.Printf("User Name %v \nNo of tickets booked %v \n", userName, userTickets)
		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)[0]
			firstNames = append(firstNames, names)
		}
		fmt.Printf("First Names of bookings %v\n", firstNames)
		fmt.Printf("Current Bookings %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Println("Sold out")
			break
		}
	}
}
