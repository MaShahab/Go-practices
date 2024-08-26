package main

import "fmt"

func main() {
	firstName := "kiarash"
	var lastName string = "amiri"

	fmt.Println(firstName)
	fmt.Println(lastName)

	currentYear := 2023
	var birthYear = 1985

	age := currentYear - birthYear

	fmt.Println(age)

	currentYear = currentYear + 1

	age = currentYear - birthYear

	fmt.Println(age)
}
