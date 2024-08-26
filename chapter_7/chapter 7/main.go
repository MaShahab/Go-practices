package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	firstName string
	lastName  string
	birthDay  string
	created   time.Time
}

func (userData *User) outputUserDetails() {
	fmt.Printf("my name is %v %v (born on %v)", userData.firstName, userData.lastName, userData.birthDay)
}

func NewUser(fName string, lName string, bDay string) *User {
	created := time.Now()

	userData := User{
		firstName: fName,
		lastName:  lName,
		birthDay:  bDay,
		created:   created,
	}

	return &userData
}

func main() {
	firstName := getUserData("please enter your first name: ")
	lastName := getUserData("please enter your last name: ")
	birthday := getUserData("please enter your birthday (MM/DD/YYYY): ")

	userData := NewUser(firstName, lastName, birthday)
	// ... do something awesome with that gathered data!

	userData.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}
