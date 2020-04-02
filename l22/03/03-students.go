package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	houses := map[string][]string{
		"gryffindor": []string{
			"weasley", "hagrid", "dumbledore", "lupin", "wenlock",
		},
		"hufflepuf": []string{
			"scamander", "helga", "diggory", "flitwick",
		},
		"ravenclaw": []string{
			"bagnold", "wildsmith", "montmorency",
		},
		"slytherin": []string{
			"horace", "nigellus", "higgs", "scorpius",
		},
		"bobo": {
			"wizardry", "unwanted",
		},
	}

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	students, ok := houses[args[0]]

	if !ok {
		fmt.Printf("Sorry. I don't know anything about %q\n", args[0])
		return
	}

	sorted := append([]string{}, students...)
	sort.Strings(sorted)

	fmt.Printf("~~~ %s students ~~~\n", args[0])

	for _, v := range sorted {
		fmt.Printf("\t + %s\n", v)
	}

}
