package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	const maxTurns = 5
	var (
		input  int
		err    error
		random int
	)

	if len(os.Args[1:]) == 0 {
		return
	}

	args := os.Args[1:]
	verbose := args[0] == "-v"
	if verbose {
		input, err = strconv.Atoi(args[1])
	} else {
		input, err = strconv.Atoi(args[0])
	}

	if err != nil {
		return
	}

	for i := 0; i < maxTurns; i++ {
		rand.Seed(time.Now().UnixNano())
		random = rand.Intn(10)
		fmt.Printf("%d ", random)
		if random == input {
			fmt.Printf("!!!YOU WIN!!!")
			break
		}
	}
	fmt.Println()
}
