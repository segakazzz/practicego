package main

import "fmt"

func main() {
	i := 314
	f := 3.14
	s := "Hello"
	b := true

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)

	a, b := 14, true
	fmt.Println(a, b)

	a, c := 42, "good"
	fmt.Println(a, c)

	sum := 27 + 3.5
	fmt.Println(sum)

	x, y := true, true
	_ = y
	fmt.Println(x)

	age, yourAge := 20, 30
	ratio, age := 1, 42
	fmt.Println(age, yourAge, ratio)
}
