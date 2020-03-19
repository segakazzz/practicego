package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = [6][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for i, a := range data {
		for _, v := range a {
			fmt.Printf("%-20s", v)
		}
		fmt.Println()
		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}
	fmt.Println(strings.Repeat("-", 50))
}
