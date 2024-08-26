package main

import "fmt"

func main() {
	radius := 5.00
	pi := 3.14

	circumference := pi * radius * 2

	fmt.Println(circumference)

	fmt.Printf("For a radius of %v, the circle circumference is %.2f", radius, circumference)
}
