package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		radius = 10.
		area   float64
	)

	// ADD YOUR CODE HERE
	area = math.Pi * radius * radius

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %g\n", radius, area)
}
