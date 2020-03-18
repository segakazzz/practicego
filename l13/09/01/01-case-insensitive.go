package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	original := "a b c d e f g h lazy"
	slice := strings.Fields(original)
	// fmt.Println(slice)
	args := os.Args[1:]
	input := strings.ToLower(args[0])
	if len(args) == 0 {
		fmt.Println("Input search keyword...")
		return
	}

	for i, w := range slice {
		// fmt.Println(i, w)
		if w == input {
			fmt.Printf("%-4d Keyword found %q\n", i, w)
			return
		}
	}

}
