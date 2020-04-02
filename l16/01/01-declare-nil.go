package main

import "fmt"

func main() {
	var friends []string
	var distances []int
	var data []byte
	var ratios []float64
	var alives []bool

	fmt.Printf("friends: %T %d %t\n", friends, len(friends), friends == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives: %T %d %t\n", alives, len(alives), alives == nil)
}
