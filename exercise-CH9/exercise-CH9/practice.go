package main

import "fmt"

type product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3)
	mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)

	// 4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5)
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics")
	fmt.Println(courseGoals)

	// 7)
	products := []product{
		{
			"first-product",
			"A first product",
			44.3,
		},
		{
			"second-product",
			"A second product",
			44.3,
		},
	}
	newProduct := product{
		"third-product",
		"A third product",
		44.3,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
