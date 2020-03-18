package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, os.Args[1]

	switch m := strings.ToLower(month); m {
	case "april", "june", "september", "november":
		days = 30
	case "january", "march", "may", "july", "august", "october", "december":
		days = 31
	case "february":
		switch {
		case leap:
			days = 29
		default:
			days = 28
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
