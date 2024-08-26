package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *product) store() {
	file, _ := os.Create(prod.id + ".txt")

	content := fmt.Sprintf(
		"ID: %v\n Title: %v\n Description: %v\n Price: %v\n",
		prod.id,
		prod.title,
		prod.description,
		prod.price,
	)

	file.WriteString(content)

	file.Close()
}

func (product *product) printData() {
	fmt.Printf("ID: %v\n", product.id)
	fmt.Printf("Title: %v\n", product.title)
	fmt.Printf("Description: %v\n", product.description)
	fmt.Printf("Price: %v\n", product.price)
}

func NewProduct(i string, t string, d string, p float64) *product {
	return &product{
		id:          i,
		title:       t,
		description: d,
		price:       p,
	}
}

func main() {
	// firstProduct := product{
	// 	"first-product",
	// 	"A book",
	// 	"A beatifull Book",
	// 	88.3,
	// }

	// secondProduct := NewProduct("second-product", "A bike", "A beatifull bike", 25.4)

	// secondProduct.printData()
	// firstProduct.printData()

	fmt.Println("Please enter product data.")
	fmt.Println("--------------------------")

	reader := bufio.NewReader(os.Stdin)

	idInput := getUserInput(reader, "Product ID: ")
	titleInput := getUserInput(reader, "Product Title: ")
	descriptionInput := getUserInput(reader, "Product Description: ")
	priceInput := getUserInput(reader, "Product Price: ")

	price, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, descriptionInput, price)

	product.printData()
	product.store()
}

func getUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
