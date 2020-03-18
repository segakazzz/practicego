package main

import (
	"fmt"
	"os"
)

func main() {
	argsCount := len(os.Args)
	if argsCount == 0 || len(os.Args[1]) != 1 {
		fmt.Println("Give me a letter")
	} else {
		letter := os.Args[1]
		if letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u" {
			fmt.Printf("%q is a vowel\n", letter)
		} else if letter == "y" || letter == "w" {
			fmt.Printf("%q is sometimes a vowel, sometimes not \n", letter)
		} else {
			fmt.Printf("%q is a consonant\n", letter)
		}
	}

}
