package main

import (
	"github.com/key-R-hash/bmi/info"
)

func main() {
	info.PrintWelcome()

	weight := getUserMetric(info.WeightPrompt)
	height := getUserMetric(info.HeightPrompt)

	bmi := calculateBMI(weight, height)

	pringBMI(bmi)
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
