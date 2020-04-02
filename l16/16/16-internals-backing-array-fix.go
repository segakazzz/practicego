package main

import "fmt"

func main() {
	nums := []int{56, 89, 15, 25, 30, 50}
	// mine := make([]int, len(nums))
	// copy(mine, nums)
	mine := append([]int{}, nums...)
	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine[:3])
	fmt.Println("Original nums:", nums[:3])
}
