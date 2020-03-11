package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "inanç           "
	name = strings.TrimRight(name, " ")
	fmt.Println(utf8.RuneCountInString(name))
	fmt.Printf("%08b\n", 1)
}
