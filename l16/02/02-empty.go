package main

import "fmt"

func main() {
	friends := []string{}
	distances := []int{}
	data := []byte{}
	ratios := []float64{}
	alives := []bool{}

	fmt.Printf("friends: %T %d %t\n", friends, len(friends), friends == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives: %T %d %t\n", alives, len(alives), alives == nil)
}
