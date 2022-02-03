package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Confrence"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking App \n", conferenceName)
	fmt.Println("We have total ", conferenceTickets, "tickets and", remainingTickets, "are avialable")
	fmt.Println("You can book tickets here")

	var userName string
	var userTickets int

	fmt.Scan(&userName)
	fmt.Scan(&userTickets)

	fmt.Printf("User Name %v \nNo of tickets booked %v \n", userName, userTickets)

}
