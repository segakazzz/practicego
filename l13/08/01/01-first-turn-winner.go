package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	guess := 6
	randNum := rand.Intn(10)
	fmt.Printf("guess %d picked %d\n", guess, randNum)
	if randNum == guess {
		fmt.Println("You Won🍀")
	}
}
