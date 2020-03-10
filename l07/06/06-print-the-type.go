package main

import "fmt"

func main() {
	value := 3
	fmt.Printf("Type of %d is %T\n", value, value)

	value2 := 3.14
	fmt.Printf("Type of %.2f is %T\n", value2, value2)

	value3 := "hello"
	fmt.Printf("Type of %s is %T\n", value3, value3)

	value4 := true
	fmt.Printf("Type of %t is %T\n", value4, value4)

}
