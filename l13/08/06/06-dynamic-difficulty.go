package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	var (
		input    int
		err      error
		random   int
		max      int
		maxTurns int
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

	if input <= 10 {
		max = 10
		maxTurns = 5
	} else if input <= 30 {
		maxTurns = 11
		max = input + 1
	} else {
		maxTurns = 30
		max = input + 1
	}

	for i := 0; i < maxTurns; i++ {
		rand.Seed(time.Now().UnixNano())
		random = rand.Intn(max)
		if verbose {
			fmt.Printf("%d ", random)
		}
		if random == input {
			fmt.Printf("!!!YOU WIN!!!")
			break
		}
	}
	fmt.Println()
}
