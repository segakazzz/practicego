package main

import (
	"bufio"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	input := map[string]bool{}
	for in.Scan() {
		line := in.Text()
		_, ok := input[line]
		switch ok {
		case true:
			return
		case false:
			input[line] = true
		}
	}
}
