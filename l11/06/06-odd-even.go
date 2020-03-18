package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Pick a number")
	} else {
		number, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("%q is not a number\n", os.Args[1])
		} else {
			isOdd := number%2 == 1
			if isOdd {
				fmt.Printf("%d is an odd number\n", number)
			} else {
				fmt.Printf("%d is an even number\n", number)
			}
		}
	}
}
