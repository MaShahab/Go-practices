package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetric(promptText string) (value float64) {
	fmt.Print(promptText)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	value, _ = strconv.ParseFloat(input, 64)

	return
}
