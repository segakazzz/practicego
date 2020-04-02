package main

import "fmt"

func main() {
	toppings := []string{"black olives", "green peppers"}
	toppings = append(toppings, "onions")
	toppings = append(toppings, "extra cheese")

	fmt.Printf("toppings       : %s\n", toppings)
}
