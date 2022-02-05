package main

import (
	validator "BookingApp/validators"
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Confrence"
var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		if !validator.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets) {
			continue
		}
		bookTickets(userTickets, firstName, lastName, email)
		sendTicket(userTickets, firstName, lastName, email)
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
		//var names = strings.Fields(booking)[0]
		//firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
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

	//var userData = make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["emaill"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	//fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets %v\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket: \n %v to email address %v\n", ticket, email)
	fmt.Println("#########################")
}
