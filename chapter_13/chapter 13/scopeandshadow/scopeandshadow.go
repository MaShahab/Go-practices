package main

import "fmt"

var userName = "Kiarash"

func main() {
	shouldContinue := true
	userName := "KiarashAmiri" // shadow the higher-level variable

	if shouldContinue {
		userAge := 30

		fmt.Printf("Name: %v, Age: %v", userName, userAge)
	}

	// fmt.Println(userAge) => Dosent't work

}

func printData() {
	fmt.Println(userName)
	// fmt.Println(shouldContinue) => Also dosent't work
}
