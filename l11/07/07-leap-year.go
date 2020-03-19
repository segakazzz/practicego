package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Give me a year number")
	} else {
		year, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("%q is not a valid year\n", os.Args[1])
		} else if year%100 == 0 || year%4 != 0 {
			fmt.Printf("%d is not a leap year.\n", year)
		} else {
			fmt.Printf("%d is a leap year.\n", year)
		}
	}
}