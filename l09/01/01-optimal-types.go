package main

import "fmt"

func main() {
	var letter byte
	letter = 'e'
	var string = "e"
	fmt.Println(letter, string)

	var unicode rune
	unicode = '佐'
	fmt.Println(unicode)

	var year uint16
	year = 2040
	fmt.Println(year)

	var month uint8
	month = 12
	fmt.Println(month)

	var speed int64
	speed = 213130129032423432
	fmt.Println(speed, int32(speed))

	var angle int16
	angle = 200
	fmt.Println(angle)
}
