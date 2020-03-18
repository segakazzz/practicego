package main

import (
	"fmt"
	"os"
)

func main() {
	argsCount := len(os.Args)

	if argsCount == 1 {
		fmt.Println("Give me args")
	} else if argsCount == 2 {
		fmt.Printf("There is one: %s\n", os.Args[1])
	} else if argsCount == 3 {
		fmt.Printf("There is two: %s %s\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("There are %d arguments \n", argsCount)
	}

}
