package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var radius, vol float64

	radius, _ = strconv.ParseFloat(os.Args[1], 64)
	// ADD YOUR CODE HERE
	vol = 4. / 3. * math.Pi * math.Pow(radius, 3.)

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
