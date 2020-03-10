package main

import (
	"fmt"
	"os"
)

func main() {
	// count := len(os.Args) - 1
	// fmt.Printf("There are %d names.\n", count)
	// args := os.Args
	// fmt.Println(args[0])

	// name := args[1]
	// fmt.Println("Hi", name)
	// fmt.Println("How are you?")

	first, second, third := os.Args[1], os.Args[2], os.Args[3]
	fmt.Println("There are", len(os.Args), "people!")
	fmt.Println("Hello great", first, "!")
	fmt.Println("Hello great", second, "!")
	fmt.Println("Hello great", third, "!")

}
