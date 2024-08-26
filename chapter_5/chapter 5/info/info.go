package info

import "fmt"

const bmiTitle = "BMI Calculator"
const seperator = "------------------"
const WeightPrompt = "Please enter your weight (kg): "
const HeightPrompt = "Please enter your height (m): "

func PrintWelcome() {
	fmt.Println(bmiTitle)
	fmt.Println(seperator)
}
