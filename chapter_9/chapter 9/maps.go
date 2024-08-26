package main

import (
	"fmt"
)

type product struct {
	id    string
	title string
	price float64
}

type websites struct {
	google   string
	aws      string
	linkedin string
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web services"])

	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
