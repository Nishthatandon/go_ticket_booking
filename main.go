package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference" //syntactic sugar format
const conferenceTickets int = 50

var remainingTickets uint = 50 // uint means that this var can not hold negative values
// declare slice type var mySlice []string
// create an empty slice
// list of string
// var bookings = []string{}
// list of maps, the last param "0" is the initial size of slice
//var bookings = make([]map[string]string, 0)
// list of struct
var bookings = make([]UserData, 0)

type UserData struct {
	firstName            string
	lastName             string
	numberOfTickets      uint
	email                string
	isOptedForNewsLetter bool
}

var wg = sync.WaitGroup{} // for thread syncronization
func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isNameValid, isEmailValid, isTicketCountValid := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isNameValid && isEmailValid && isTicketCountValid {

			bookTickets(userTickets, firstName, lastName, email)
			// go keyword makes the process async/concurrent
			wg.Add(1) // set the numbner of go routines the main method has to wait for before exiting
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isNameValid {
				fmt.Println("Either of the first or last name entered is too short")
			}
			if !isEmailValid {
				fmt.Println("Email address is not valid.")
			}
			if !isTicketCountValid {
				fmt.Println("The requestes number of tickets is invalid.")
			}
		}
	}
	wg.Wait() // blocks until the wait group counter is 0
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// string splitting with Fields
		//firstNames = append(firstNames, strings.Fields(booking)[0])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// map type declaration => var userData map[string]string
	//create an empty map
	// var userData = make(map[string]string)
	//assign values
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// type conversion
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName:            firstName,
		lastName:             lastName,
		numberOfTickets:      userTickets,
		email:                email,
		isOptedForNewsLetter: false,
	}
	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a booking confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets, for %v %v", firstName, lastName, userTickets)
	fmt.Println("***************")
	fmt.Printf("sending ticket: \n %v \nto email %v\n", ticket, email)
	fmt.Println("***************")
	wg.Done()
}
