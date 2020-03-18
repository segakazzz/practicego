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
	)

	if len(os.Args) < 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	if size, err = strconv.ParseInt(os.Args[1], 10, 64); err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i < int(size); i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Printf("\n")
	for i := 0; i < int(size); i++ {
		fmt.Printf("%5d", i)
		for j := 0; j < int(size); j++ {
			fmt.Printf("%5d", i+j)
		}
		fmt.Printf("\n")
	}

}
