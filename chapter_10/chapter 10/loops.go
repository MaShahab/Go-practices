package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice, exiting!")
		return
	}

	if selectedChoice == "1" {
		calculateSumUpToNumber()
	} else if selectedChoice == "2" {
		calculateFactorial()
	} else if selectedChoice == "3" {
		calculateSumManually()
	} else {
		calculateListSum()
	}
}

func calculateSumUpToNumber() {
	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid Number input")
		return
	}

	fmt.Println(chosenNumber)

	sum := 0

	for i := 1; i <= chosenNumber; i++ {
		sum = sum + i // 5 + 4 + 3 + 2 + 1 = 15
	}

	fmt.Printf("Result: %v", sum)
}

func calculateFactorial() {
	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid Number input")
		return
	}

	factorial := 1

	for i := 1; i <= chosenNumber; i++ {
		factorial = factorial * i // 5 * 4 * 3 * 2 * 1 = 120
	}

	fmt.Printf("Result: %v", factorial)
}

func calculateSumManually() {
	isEnteringNumbers := true
	sum := 0

	fmt.Println("Keep on entering numbers, the proggram will quite one you enter any other value.")

	for isEnteringNumbers {
		chosenNumber, err := getInputNumber()
		sum = sum + chosenNumber

		isEnteringNumbers = err == nil
	}

	fmt.Printf("Result: %v\n", sum)
}

func calculateListSum() {
	fmt.Println("Please enter a comma-seperated list of numbers.")

	inputNumberList, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	inputNumberList = strings.Replace(inputNumberList, "\n", "", -1)
	inputNumbers := strings.Split(inputNumberList, ",")

	sum := 0

	for index, value := range inputNumbers {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
		number, err := strconv.ParseInt(value, 0, 64)

		if err != nil {
			continue
		}

		sum = sum + int(number)
	}

	fmt.Printf("Result %v\n", sum)
}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	choseNumber, err := strconv.ParseInt(inputNumber, 0, 64)

	if err != nil {
		return 0, err
	}

	return int(choseNumber), nil
}

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice")
	fmt.Println("1) Add up all the number od to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually enterned numbers")
	fmt.Println("4) Sum up a list of enterned numbers")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("Invalid input!")
	}
}
