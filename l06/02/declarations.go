package main

import "fmt"

var varpackagescope int

func main() {
	// int
	var newvar int
	fmt.Println(newvar)

	// bool
	var isOn bool
	fmt.Println(isOn)

	// float64
	var brightness float64
	fmt.Println(brightness)

	// string
	var stringvar string
	fmt.Printf("stringvar (%T): %q\n", stringvar, stringvar)

	// undeclarables
	// var (
	// 	3speed int
	// 	!speed int
	// 	spe?ed int
	// 	var int
	// 	func int
	// 	package int
	// )

	// with bits
	var (
		a int
		b int8
		c int16
		d int32
		e int64
		f float32
		g float64
		h complex64
		i complex128
		j bool
		k string
		l rune
		m byte
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m)

	var (
		active bool
		delta  int
	)
	fmt.Println(active, delta)

	var firstName, lastName string
	fmt.Printf("%q, %q\n", firstName, lastName)

	var isLiquid bool
	_ = isLiquid

	var valnotexistnow int
	fmt.Println(valnotexistnow)
}
