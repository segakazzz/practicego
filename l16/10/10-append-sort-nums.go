package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		number int
		err    error
	)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	inputs := []int{}

	for _, v := range args {
		number, err = strconv.Atoi(v)
		// fmt.Println(number)
		if err == nil {
			inputs = append(inputs, number)
		}
	}

	fmt.Printf("%v\n", inputs)
}
