package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := generateRandomNumber()

	sum := add(a, b)
	printNumber(sum)
}

func generateRandomNumber() (r1 int, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(10)

	return
}

func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

func printNumber(num int) {
	fmt.Printf("this is my number %v", num)
}
