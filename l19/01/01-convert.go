package main

import "fmt"

func main() {

	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	for _, word := range words {
		bwords = append(bwords, []byte(word))
	}

	for _, bword := range bwords {
		fmt.Printf("% x, %s\n", bword, string(bword))
	}
}
