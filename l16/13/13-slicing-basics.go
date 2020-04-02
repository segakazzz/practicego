package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"
	nums := []int{}

	for _, v := range strings.Fields(data) {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}

	fmt.Printf("nums: %[1]d %[1]T\n", nums)

	fmt.Printf("evens: %[1]d %[1]T\n", nums[0:3])
	fmt.Printf("odds: %[1]d %[1]T\n", nums[3:6])
	fmt.Printf("middle: %[1]d %[1]T\n", nums[2:4])
	fmt.Printf("first two: %[1]d %[1]T\n", nums[0:2])
	fmt.Printf("last two: %[1]d %[1]T\n", nums[len(nums)-2:])

	evens := nums[:3]
	odds := nums[3:]

	fmt.Printf("even last: %[1]d %[1]T\n", evens[len(evens)-1:])
	fmt.Printf("odd last: %[1]d %[1]T\n", odds[len(odds)-2:])

}
