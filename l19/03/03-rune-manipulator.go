package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, str := range strings {
		fmt.Printf("%-20s bytelen: %d runelen: %d\n", str, len(str), utf8.RuneCountInString(str))
		fmt.Printf("% x\n", str)

		for _, r := range str {
			fmt.Printf("%x ", r)
		}
		fmt.Println()

		for _, r := range str {
			fmt.Printf("%s ", string(r))
		}
		fmt.Println()
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%s size:%d\n", string(r), size)
		r2, size2 := utf8.DecodeRuneInString(str[size:])
		fmt.Printf("%s size:%d\n", string(r2), size2)

		l, sizel := utf8.DecodeLastRuneInString(str)
		fmt.Printf("%s size:%d\n", string(l), sizel)

		l2, sizel2 := utf8.DecodeLastRuneInString(str[:len(str)-sizel])
		fmt.Printf("%s size:%d\n", string(l2), sizel2)

		convert := []rune(str)
		fmt.Printf("%s %s\n", string(convert[0]), string(convert[1]))
		fmt.Printf("%s %s\n", string(convert[len(convert)-2]), string(convert[len(convert)-1]))
	}

}
