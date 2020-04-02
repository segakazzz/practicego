package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	p := newParser()

	var err error
	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p, err = parseField(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Print the visits per domain
	printVisits(p)

	// Print the total visits for all domains
	printTotalVisits(p)

	// Let's handle the error
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
