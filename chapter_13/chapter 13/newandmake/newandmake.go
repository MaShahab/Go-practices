package main

import "fmt"

func main() {
	// new & make
	number := new(int)
	fmt.Println(number)
	fmt.Println(*number)

	anotherNumber := 0
	numberAdrress := &anotherNumber

	fmt.Println(anotherNumber)
	fmt.Println(numberAdrress)

	// hobbies := []string{"Sports", "Reading"}
	hobbies := make([]string, 2, 10)
	// aMap := make(map[string]string, 5)
	moreHobbies := new([]string)

	fmt.Println(*moreHobbies)

	*moreHobbies = append(*moreHobbies, "Sports")

	fmt.Println(*moreHobbies)

	hobbies[0] = "Sports"
	hobbies[1] = "Reading"
	// hobbies[2] = "Cooking2"

	fmt.Println(hobbies)

	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)
}
