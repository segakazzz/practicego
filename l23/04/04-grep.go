package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		if strings.Contains(line, args[0]) {
			fmt.Println(line)
		}
	}
}
