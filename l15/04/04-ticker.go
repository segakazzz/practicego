package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec, unix := now.Hour(), now.Minute(), now.Second(), now.Unix()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}
		newClock := clock
		startIndex := int(unix) % 16
		if startIndex < 8 {
			for i := range clock {
				if startIndex+i >= len(clock) {
					newClock[i] = empty
				} else {
					newClock[i] = clock[startIndex+i]
				}
			}
		} else {
			numOfEmpty := 16 - startIndex
			for i := range clock {
				if i < numOfEmpty {
					newClock[i] = empty
				} else {
					newClock[i] = clock[i-numOfEmpty]
				}
			}
		}

		for line := range newClock[0] {
			for index, digit := range newClock {
				// colon blink
				next := newClock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
