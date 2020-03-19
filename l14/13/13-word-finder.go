package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please give me a word to search.")
		return
	}

	words := strings.Fields(corpus)
	filter := [...]string{
		"and",
		"or",
		"was",
		"the",
		"since",
		"very",
	}

label:
	for _, a := range args {
		// Check if the word is in filter
		for _, f := range filter {
			if strings.ToLower(a) == strings.ToLower(f) {
				continue label
			}
		}

		for i, w := range words {
			if strings.ToLower(a) == strings.ToLower(w) {
				fmt.Printf("#%-5d: %q\n", i+1, w)
			}
		}
	}

}
