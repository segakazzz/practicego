package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		size int64
		err  error
		op   string
	)

	if len(os.Args) < 3 {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	if size, err = strconv.ParseInt(os.Args[2], 10, 64); err != nil || size <= 0 {
		fmt.Println("Size is missing")
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	switch os.Args[1] {
	case "*", "+", "-", "/", "%":
		op = os.Args[1]
	default:
		fmt.Println("Invalid operator.")
		fmt.Println("Valid ops one of: */+-")
		return
	}

	fmt.Printf("%5s", op)
	for i := 0; i < int(size); i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Printf("\n")
	for i := 0; i < int(size); i++ {
		fmt.Printf("%5d", i)
		for j := 0; j < int(size); j++ {
			switch op {
			case "+":
				fmt.Printf("%5d", i+j)
			case "-":
				fmt.Printf("%5d", i-j)
			case "*":
				fmt.Printf("%5d", i*j)
			case "/":
				fmt.Printf("%5d", i/j)
			case "%":
				if i == 0 || j == 0 {
					fmt.Printf("%5d", 0)
				} else {
					fmt.Printf("%5d", i%j)
				}
			}
		}
		fmt.Printf("\n")
	}

}
