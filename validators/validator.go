package validator

import (
	"fmt"
	"strings"
)

func ValidateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	isValidUserInput := true
	if !isValidName {
		fmt.Println("Invalid Name")
		isValidUserInput = false
	}
	if !isValidEmail {
		fmt.Println("Invalid Email")
		isValidUserInput = false
	}
	if !isValidTicket {
		fmt.Println("Invalid Ticket")
		isValidUserInput = false
	}
	return isValidUserInput
}
