package main

import "fmt"

type computer struct {
	brand string
}

func main() {

	// create a nil pointer with the type of pointer to a computer
	var P *computer
	// compare the pointer variable to a nil value, and say it's nil
	fmt.Printf("pointer: %p nil?: %t\n", P, P == nil)

	// create an apple computer by putting its address to a pointer variable
	apple := computer{brand: "apple"}
	P = &apple
	fmt.Printf("pointer: %p nil?: %t value: %#v \n", P, P == nil, *P)

	// put the apple into a new pointer variable
	P2 := &apple

	// compare the apples: if they are equal say so and print their addresses
	if *P == *P2 {
		fmt.Printf("pointer1: %p, pointer2: %p, value: %#v \n", P, P2, *P)
	}

	// create a sony computer by putting its address to a new pointer variable
	*P = computer{brand: "sony"}
	fmt.Printf("address %p value: %#v\n", P, *P)

	// compare apple to sony, if they are not equal say so and print their
	// addresses

	// if *P == *P2 {
	fmt.Printf("pointer1: %p, pointer2: %p, value: %#v, value2: %#v \n", P, P2, *P, *P2)
	// }

	// put apple's value into a new ordinary variable

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so

	// print the values:
	// the value that is pointed by the apple and the new variable

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
}
