package main

import (
	"fmt"
	"path"
)

func main() {

	color := "green"
	color = "blue"
	fmt.Println("color >>> ", color)

	color02 := "green"
	color02 = "dark " + color02
	fmt.Println("color02 >>> ", color02)

	n := 3.14
	n *= 2
	fmt.Println("n >>> ", n)

	var perimeter int
	width := 5
	height := 6
	perimeter = (width + height) * 2
	fmt.Println("perimeter >>> ", perimeter)

	var (
		lang    string
		version int
	)
	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on ", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	a, b := multi()
	_ = a
	fmt.Println(b)

	color1, color2 := "red", "blue"
	color1, color2 = "orange", "green"
	fmt.Println(color1, color2)

	red, blue := "red", "blue"
	blue, red = red, blue
	fmt.Println(red, blue)

	dir, _ := path.Split("secret/file.txt")
	fmt.Println(dir)
}

func multi() (int, int) {
	return 5, 4
}
