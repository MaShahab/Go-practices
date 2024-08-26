package main

import (
	"fmt"
)

func main() {
	age := 30

	fmt.Println(age)

	var myage *int
	myage = &age

	fmt.Println(*myage)

	*myage = 33

	fmt.Println(*myage)
	fmt.Println(age)

	doubledAge := double(&age)
	fmt.Println(doubledAge)
	fmt.Println(age)
}

func double(number *int) int {
	result := *number * 2
	*number = 100
	return result
}
