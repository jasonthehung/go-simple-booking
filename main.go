package main

import (
	"fmt"
	"simple-booking/helper"
	"sync"
	"time"
)

const conferenceTickets uint = 50

// -- Package Level Variable -- /
// ** Use in the same package ** //
var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []UserData // slice

// Create a custom tpye called "UserData" based on a struct {...}
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Creare a wait group
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		// Use a new thread to execute this function
		go sendTicket(userTickets, firstName, lastName, email)

		// call function to get first name
		fmt.Printf("The first names of bookings are: %v\n", getFirstNameList())

		if remainingTickets == 0 {
			// end program
			fmt.Printf("Our conference is booked out. Come back next year.")
			// break
		}
	} else if userTickets == remainingTickets {
		// do something else...
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid.")
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking applicatrion.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")
}

func getFirstNameList() []string {
	var firstNames []string

	// Blank Identifier (To ignore var we don't want to use)
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at the end of the day.\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", conferenceTickets, remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("###############")

	wg.Done()
}
