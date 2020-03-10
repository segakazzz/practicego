package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Your name is %s and your lastname is %s\n", args[1], args[2])
}
