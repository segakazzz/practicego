package main

import (
	"fmt"
	"strings"
)

func main() {
	var books = [3]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
	}

	upper := books
	lower := books

	for i, v := range upper {
		upper[i] = strings.ToUpper(v)
	}
	for i, v := range lower {
		lower[i] = strings.ToLower(v)
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)

}
