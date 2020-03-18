package main

import "fmt"

func main() {
	const minsPerDay = 60 * 24
	const weekDays = 7
	fmt.Printf("Total Number of minutest in two weeks is %d\n", minsPerDay*weekDays*2)
}
