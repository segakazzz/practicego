package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := strings.Split("Da Vinci, Wozniak, Carmack", ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	sort.Strings(namesA)
	sort.Strings(namesB)

	fmt.Printf("%v\n", namesA)
	fmt.Printf("%v\n", namesB)

	for i := range namesA {
		if namesA[i] != namesB[i] {
			break
		}
		if i == len(namesA)-1 {
			fmt.Println("They are equal.")
		}
	}

}
