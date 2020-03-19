package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	const (
		EUR = iota
		GBP
		JPY
	)

	var ratio = [...]float64{
		EUR: 0.92,
		GBP: 0.87,
		JPY: 108.88,
	}

	var key = [...]string{"EUR", "GBP", "JPY"}

	if len(args) == 0 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	given, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
	}

	for i, v := range ratio {
		fmt.Printf("%.2f USD is %.2f %s\n", given, given*v, key[i])
	}

}
