package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// You need to add a newline after each sentence in another slice.
	// Don't touch the following code.
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday`)

	// ===================================
	//
	// ~~~ CHANGE THIS CODE ~~~
	//
	// fix := lyric
	fix := make([]string, len(lyric)+2)
	copy(fix[0:8], lyric[0:8])
	fix[8] = "\n"
	copy(fix[9:19], lyric[8:18])
	fix[19] = "\n"
	copy(fix[20:], lyric[18:])
	//
	// ===================================

	// Currently, it prints every sentence on the same line.
	// Don't touch the following code.

	s.Show("fix slice", lyric)
	s.Show("fix slice", fix)

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}

func init() {
	//
	// YOU DON'T NEED TO TOUCH THIS
	//
	// This initializes some options for the prettyslice package.
	// You can change the options if you want.
	//
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	//
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 5       // prints max 15 elements per line
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}
