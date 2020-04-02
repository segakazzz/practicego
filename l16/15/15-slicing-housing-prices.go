package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		dataRow    []string
		start, end int
	)

	dataRows := strings.Split(data, "\n")
	titleRow := strings.Split(dataRows[0], ",")
	end = len(titleRow)
	args := os.Args[1:]

	switch len(args) {
	case 1:
		for i, title := range titleRow {
			if title == args[0] {
				start = i
				break
			}
		}
	case 2:
		for i, title := range titleRow {
			if title == args[0] {
				start = i
				break
			}
		}
		for i, title := range titleRow {
			if title == args[1] {
				end = i
				break
			}
		}
	}

	if start > end {
		start = 0
	}

	for _, row := range dataRows {
		dataRow = strings.Split(row, ",")
		for i, v := range dataRow {
			if i >= start && i <= end {
				fmt.Printf("%10s", v)
			}
		}
		fmt.Println()
	}

}
