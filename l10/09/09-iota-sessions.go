package main

import "fmt"

func main() {

	const (
		Winter = 12 - (iota * 3)
		Fall
		Summer
		Spring
	)

	fmt.Println(Winter, Fall, Summer, Spring)

	// fmt.Println(int(true))
}
