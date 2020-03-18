package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	if len(os.Args[1:]) < 2 {
		fmt.Println("Input 2 numbers after command....")
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	guess2, err2 := strconv.Atoi(os.Args[2])

	if err != nil || guess <= 0 {
		fmt.Printf("%q is not a valid number\n", os.Args[1])
		return
	}
	if err2 != nil || guess2 <= 0 {
		fmt.Printf("%q is not a valid number\n", os.Args[2])
		return
	}

	randNum := rand.Intn(10)
	fmt.Printf("guess %d and %d picked %d\n", guess, guess2, randNum)
	randMessage := rand.Intn(2)
	if randNum == guess || randNum == guess2 {
		switch randMessage {
		case 0:
			fmt.Println("YOU WON")
		case 1:
			fmt.Println("YOU'RE AWESOME")
		}
	} else {
		switch randMessage {
		case 0:
			fmt.Println("LOSER!")
		case 1:
			fmt.Println("YOU LOST. TRY AGAIN?")
		}
	}

}
