package main
import "fmt"

func main(){
	conferenceName := "Go conference" // constant not allowed
	const conferenceTickets int= 50
	var remainingTickets int= 50
	
	fmt.Printf("%T\n",conferenceTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, %v are still available\n", conferenceTickets,remainingTickets)

	var userName string
	var userTicket int
	// ask user for name
	fmt.Scan(userName)

	userName = "Tom"
	userTicket = 2
	fmt.Printf("User %v booked %v tickets\n",userName,userTicket)
}