package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var books = [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Tell me a book title")
		return
	}

	matched := 0

	fmt.Println("Search Results:")
	for _, v := range books {
		if strings.Contains(strings.ToLower(v), strings.ToLower(args[0])) {
			fmt.Printf("+ %s\n", v)
			matched++
		}
	}
	if matched == 0 {
		fmt.Printf("We don't have the book: %q\n", args[0])
	}
}
