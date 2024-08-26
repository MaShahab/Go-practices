package lists

import "fmt"

func main() {
	prices := []float64{33.1, 77.2}
	fmt.Println(prices[0:1])
	prices[1] = 12.2

	prices = append(prices, 3.9)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{66.4, 33.2, 12.1}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{
// 		"book",
// 	}
// 	prices := [4]float64{44.3, 55.1, 2.4, 99.7}
// 	fmt.Println(prices[2])

// 	productNames[2] = "bike"

// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.8
// 	highlighttedPrices := featuredPrices[:1]
// 	fmt.Println(highlighttedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlighttedPrices), cap(highlighttedPrices))

// 	highlighttedPrices = highlighttedPrices[:3]
// 	fmt.Println(len(highlighttedPrices), cap(highlighttedPrices))
// 	fmt.Println(highlighttedPrices)
// }
