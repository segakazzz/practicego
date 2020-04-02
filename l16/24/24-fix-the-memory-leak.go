package main

import (
	"fmt"
	"io/ioutil"

	"github.com/segakazzz/learngo/16-slices/exercises/24-fix-the-memory-leak/api"
)

func main() {
	// reports the initial memory usage
	api.Report()

	// returns a slice with 10 million elements.
	// it allocates 65 MB of memory space.

	millions := api.Read()

	// -----------------------------------------------------
	// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪

	// last10 := millions[len(millions)-10:]
	// last10 := append([]int{}, millions[len(millions)-10:]...)

	last10 := make([]int, 10)
	copy(last10, millions[len(millions)-10:])
	millions = last10

	fmt.Printf("\nLast 10 elements: %d\n\n", last10)

	// ✪ ONLY CHANGE THE CODE IN THIS AREA ✪
	// -----------------------------------------------------

	api.Report()

	// don't worry about this code.
	fmt.Fprintln(ioutil.Discard, millions[0])
}
