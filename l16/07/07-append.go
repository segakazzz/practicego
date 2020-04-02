package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 1. uncomment the code below
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 2. append elements to header to make it equal with the png slice
	header = append(header, 'P')
	header = append(header, 'N')
	header = append(header, 'G')

	fmt.Printf("%t\n", bytes.Equal(header, png))
	// 3. compare the slices using the bytes.Equal function

	// 4. print whether they're equal or not
}
