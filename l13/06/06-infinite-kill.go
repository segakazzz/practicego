package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	const milliseconds = 250
	var chars = strings.Fields("/ - \\ |")

	for {
		fmt.Printf("\r%s Please Wait. Processing....\n", chars[rand.Intn(4)])
		// fmt.Println(rand.Intn(4))
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
	}
}
