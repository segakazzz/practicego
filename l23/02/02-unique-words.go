package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(os.Stdin)
	words := map[string]int{}
	wordsList := []string{}

	in.Split(bufio.ScanWords)
	for in.Scan() {
		w := in.Text()
		_, isExists := words[w]
		if isExists {
			words[w]++
		} else {
			wordsList = append(wordsList, w)
			words[w] = 1
		}
	}

	for _, v := range wordsList {
		fmt.Printf("%-20s %d\n", v, words[v])
	}
}
