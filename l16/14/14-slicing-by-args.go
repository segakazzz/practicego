package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	args := os.Args[1:]

	var (
		start, end       int
		errStart, errEnd error
	)

	if len(args) != 2 && len(args) != 1 {
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	}

	switch len(args) {
	case 1:
		start, errStart = strconv.Atoi(args[0])
		if errStart != nil || start < 0 {
			fmt.Println("Wrong Positions")
			return
		}
		fmt.Printf("%s\n", ships[start:])

	case 2:
		start, errStart = strconv.Atoi(args[0])
		end, errEnd = strconv.Atoi(args[1])

		if errStart != nil || errEnd != nil || start < 0 || end > len(ships) || start > end {
			fmt.Println("Wrong Positions")
			return
		}
		fmt.Printf("%s\n", ships[start:end])
	}

}
