package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Give me a month name")
		return
	}

	month := strings.ToLower(os.Args[1])
	if month == "january" || month == "march" || month == "may" || month == "july" || month == "august" || month == "october" || month == "december" {
		fmt.Printf("%q has 31 days.\n", month)
	} else if month == "february" {
		fmt.Printf("%q has 28 days.\n", month)
	} else if month == "april" || month == "june" || month == "september" || month == "november" {
		fmt.Printf("%q has 30 days.\n", month)
	} else {
		fmt.Printf("%q is not a month.\n", month)
	}

}
