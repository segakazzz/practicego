package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	in := bufio.NewScanner(os.Stdin)
	words := map[string]int{}
	wordsList := []string{}

	in.Split(bufio.ScanWords)
	for in.Scan() {
		w := in.Text()
		single := []rune{}
		for _, v := range w {
			if !unicode.IsPunct(v) {
				single = append(single, v)
			}
		}
		w = strings.ToLower(string(single))
		_, isExists := words[w]
		if isExists {
			words[w]++
		} else {
			wordsList = append(wordsList, w)
			words[w] = 1
		}
	}

	for i, v := range wordsList {
		fmt.Printf("%10d %-20s %d\n", i+1, v, words[v])
	}
}
