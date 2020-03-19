package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Kazue"
	names[1] = "Yosuke"
	names[2] = "Hana"
	fmt.Printf("names: 	%#v\n", names)
	fmt.Printf("names: 	%q\n", names)
	fmt.Printf("names: 	%T\n", names)
	var distances [5]int
	distances[0] = 43
	distances[1] = 44
	distances[2] = 45
	distances[3] = 46
	distances[4] = 47
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("distances: %T\n", distances)
	var data [5]byte
	data[0] = 1
	data[1] = 2
	data[2] = 3
	data[3] = 4
	data[4] = 5
	fmt.Printf("data: %#v\n", data)
	fmt.Printf("data: %d\n", data)
	fmt.Printf("data: %T\n", data)
	var ratios [1]float64
	ratios[0] = 5.2
	fmt.Printf("ratios: %#v\n", ratios)
	fmt.Printf("ratios: %f\n", ratios)
	fmt.Printf("ratios: %T\n", ratios)
	var alives [2]bool
	alives[0] = true
	alives[1] = false
	fmt.Printf("alives: %#v\n", alives)
	fmt.Printf("alives: %b\n", alives)
	fmt.Printf("alives: %T\n", alives)
	var zero [0]byte
	fmt.Printf("zero: %#v\n", zero)
	fmt.Printf("zero: %b\n", zero)
	fmt.Printf("zero: %T\n", zero)
}
