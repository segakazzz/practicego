package main

import "fmt"

func main() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	fmt.Println(5.5)

	age := 2
	fmt.Println(7.5 + float64(age))

	min := int8(127)
	max := int16(1000)
	fmt.Println(max + int16(min))
}
