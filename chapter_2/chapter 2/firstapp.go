package main

import (
	"fmt"

	"github.com/yourorg/firstapp/greeting"
)

func main() {
	// var greetingText string
	// greetingText = "Hello world!!"

	// var greetingText string = "Hello world!!"
	luckyNumber := 13
	var evenMoreLuckyNumber int = luckyNumber + 5

	fmt.Println(greeting.GreetingText)
	fmt.Println(greeting.GreetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3

	fmt.Println(newNumber)

	defaultFloat := 2.123456789123456789123456789
	var smallFloat float32 = 2.123456789123456789123456789

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var newRune rune = '$'
	fmt.Println(newRune)
	fmt.Println(string(newRune))

	var newByte byte = 'a'
	fmt.Println(newByte)
	fmt.Println(string(newByte))

	firstName := "kiarash"
	lastName := "amiri"

	// someNumber := "10"
	// fmt.Println(int(someNumber) + 9)

	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 32

	fmt.Printf("Hi, my full name is %v and my age is %v (Type: %T) \n", fullName, age, age)

	fmt.Println(pi)
}
