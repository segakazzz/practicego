package main

import "fmt"

//
//  1. Declare and print the following arrays with their types:
//
//    1. The names of your best three friends (names array)
//
//    2. The distances to five different locations (distances array)
//
//    3. A data buffer with five bytes of capacity (data array)
//
//    4. Currency exchange ratios only for a single currency (ratios array)
//
//    5. Up/Down status of four different web servers (alives array)
//
//    6. A byte array that doesn't occupy memory space (zero array)
//
//  2. Print only the types of the same arrays.
//
//  3. Print only the elements of the same arrays.

func main() {
	var names [3]string
	fmt.Printf("names: 	%#v\n", names)
	fmt.Printf("names: 	%q\n", names)
	fmt.Printf("names: 	%T\n", names)
	var distances [5]int
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("distances: %T\n", distances)
	var data [5]byte
	fmt.Printf("data: %#v\n", data)
	fmt.Printf("data: %d\n", data)
	fmt.Printf("data: %T\n", data)
	var ratios [1]float64
	fmt.Printf("ratios: %#v\n", ratios)
	fmt.Printf("ratios: %f\n", ratios)
	fmt.Printf("ratios: %T\n", ratios)
	var alives [2]bool
	fmt.Printf("alives: %#v\n", alives)
	fmt.Printf("alives: %b\n", alives)
	fmt.Printf("alives: %T\n", alives)
	var zero [0]byte
	fmt.Printf("zero: %#v\n", zero)
	fmt.Printf("zero: %d\n", zero)
	fmt.Printf("zero: %T\n", zero)
}
