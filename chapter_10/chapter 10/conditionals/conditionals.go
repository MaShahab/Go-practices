package conditionals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	userAge, err := getUserInput()

	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	// isOldEnough := userAge >= 18

	if (userAge >= 30 && userAge < 50) || userAge >= 60 {
		fmt.Println("you are eligible for our senior jobs.")
	} else if userAge >= 50 {
		fmt.Println("The best age!")
	} else if userAge >= 18 {
		fmt.Println("Welcome to the club!")
	} else {
		fmt.Println("Sorry, you are not old enough!")
	}

	fmt.Println("Goodbye!")
}

func getUserInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

	return int(userAge), err
}
