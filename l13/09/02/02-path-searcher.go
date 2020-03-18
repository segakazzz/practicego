package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Getenv("PATH")
	search := strings.Split(path, ":")
	var dirs []string
	for i, p := range search {
		fmt.Printf("%2d %s", i, p)
		dirs = strings.Split(path, "/")
		for _, d := range dirs {
			if os.Args[1] == d {
				fmt.Printf(" !!! MATCH !!!")
			}
		}
		fmt.Println()
	}
}
