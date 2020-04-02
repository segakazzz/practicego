package main

import (
	"fmt"
	"strings"
)

func main() {
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)
	// newLyric := []string{"yesterday"}
	lyric = append([]string{"yesterday"}, lyric...)
	// newLyric[0] = "yesterday"
	// for _, v := range lyric {
	// 	fmt.Println(v)
	// }

	// fmt.Println(newLyric)
	fmt.Println(lyric)
}
