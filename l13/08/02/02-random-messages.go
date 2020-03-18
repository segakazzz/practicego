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
	if len(os.Args[1:]) == 0 {
		fmt.Println("Input the number after command....")
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess <= 0 {
		fmt.Printf("%q is not a valid number\n", os.Args[1])
		return
	}

	randNum := rand.Intn(10)
	fmt.Printf("guess %d picked %d\n", guess, randNum)
	randMessage := rand.Intn(2)
	if randNum == guess {
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
