package main

import "fmt"

func main() {
	friends := []string{"a", "b", "c"}
	distances := []int{0, 2, 3, 4}
	data := []byte{6, 2, 7, 8}
	ratios := []float64{0.5, .4, 2.4, .54}
	alives := []bool{true, true, false, false}

	fmt.Printf("friends: %T %d %t\n", friends, len(friends), friends == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives: %T %d %t\n", alives, len(alives), alives == nil)
}
