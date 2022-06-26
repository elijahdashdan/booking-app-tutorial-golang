package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go conference" // constant not allowed
	const conferenceTickets int= 50
	var remainingTickets uint= 50
	
	fmt.Printf("%T\n",conferenceTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, %v are still available\n", conferenceTickets,remainingTickets)

	// var bookings = [50] string{}
	var bookings []string

	for {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	// ask user for name
	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter no. of ticket")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Println(len(bookings))
	fmt.Println(bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. Please see confirmation email at %v\n",firstName,lastName,userTicket,email)
	fmt.Printf("%v remaining ticket\n",remainingTickets)

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames,names[0])
	}
	fmt.Printf("The first names are %v\n", firstNames)
	}
	
}