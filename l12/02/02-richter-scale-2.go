package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Tell me the magnitude of earthquack in human terms.")
		return
	}

	term := os.Args[1]

	switch term {
	case "micro":
		fmt.Printf("%s's richter scale is less than 2.0\n", term)
	case "very minor":
		fmt.Printf("%s's richter scale is 2 - 2.9\n", term)
	case "minor":
		fmt.Printf("%s's richter scale is 3 - 3.9\n", term)
	case "light":
		fmt.Printf("%s's richter scale is 4 - 4.9\n", term)
	case "moderate":
		fmt.Printf("%s's richter scale is 5 - 5.9\n", term)
	case "strong":
		fmt.Printf("%s's richter scale is 6 - 6.9\n", term)
	case "major":
		fmt.Printf("%s's richter scale is 7 - 7.9\n", term)
	case "great":
		fmt.Printf("%s's richter scale is 8 - 9.9\n", term)
	case "massive":
		fmt.Printf("%s's richter scale is 10+\n", term)
	default:
		fmt.Printf("%s's richter scale is unknown\n", term)
	}

}
