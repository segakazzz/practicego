package main

import "fmt"

const loopCount = 10e6

func main() {
	var slice = []string{}
	var prevCap = 0

	for i := 0; i < loopCount; i++ {
		slice = append(slice, "")

		if prevCap != cap(slice) {
			fmt.Printf("len: %d  cap: %d ratio:%.2f\n", len(slice), cap(slice), float64(cap(slice))/float64(prevCap))
		}
		prevCap = cap(slice)
	}

}
