package main

import "fmt"

func main() {
	numbers := []int{1, 10, 20}
	sum := sumup(4, 5, 10, 20, 11, 8)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}
