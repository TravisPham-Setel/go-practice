package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50 //can not re declare
var remainTickets int = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := helper.GetUserInput()
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")

		if userTickets > remainTickets {
			fmt.Println("Fuck you scammer")
			break
		}

		if isValidEmail && isValidName {
			bookings = append(bookings, firstName+" "+lastName)
			firstNames := getFirstNames(bookings)

			remainTickets = remainTickets - userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("the first name of bookings is: %v\n", firstNames)
			fmt.Printf("these are all our bookings: %v\n", bookings)
			fmt.Printf("%v tickets remain\n", remainTickets)

			if remainTickets == 0 {
				fmt.Println("Our tickets sold out, bye!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("name too short")
			}
			if !isValidEmail {
				fmt.Println("email not match convention")
			}
			fmt.Println("inValid input try again")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainTickets, "are still available")
	fmt.Println("get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}
