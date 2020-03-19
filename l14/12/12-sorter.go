package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n1   int
		n2   int
		err1 error
		err2 error
	)
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 5 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	}

	for i, v1 := range args {
		n1, err1 = strconv.Atoi(v1)
		for j, v2 := range args {
			n2, err2 = strconv.Atoi(v2)
			if i != j && err1 == nil && err2 == nil && n1 < n2 {
				args[i], args[j] = args[j], args[i]
				fmt.Println(args)
			}
		}
	}
}
