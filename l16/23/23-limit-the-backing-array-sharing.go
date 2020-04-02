package main

import (
	"fmt"

	"github.com/segakazzz/learngo/16-slices/exercises/23-limit-the-backing-array-sharing/api"
)

func main() {
	// DO NOT CHANGE ANYTHING IN THIS CODE.
	fmt.Println("(original)api.temps:", api.All())

	// get the first three elements from api.temps
	received := append([]int{}, api.Read(0, 3)...)

	// append changes the api package's temps slice's
	// backing array as well.
	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.received :", received)
}
