package main

import (
	"fmt"
)

type transformFn func(int) int

// type anotherFn func(string, int, []string, map[string]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{3, 5, 2}
	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformedFn1 := getTranformedFunction(&numbers)
	transformedFn2 := getTranformedFunction(&moreNumbers)

	transformedNumbers := transformNumber(&numbers, transformedFn1)
	moreTranformedNumbers := transformNumber(&moreNumbers, transformedFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTranformedNumbers)
}

func transformNumber(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func getTranformedFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}

}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
