package main

import (
	"fmt"
	"time"
)

func main() {
	pizza := []string{
		"onion",
		"sausage",
		"bacon",
	}

	pizza = append(pizza, "tomato")

	graduations := []int{2013, 2009, 2020}
	graduations = append(graduations, 1999)

	now := time.Now()
	departures := []time.Time{
		now,
		now.Add(time.Hour * 1),
		now.Add(time.Hour * 2),
		now.Add(time.Hour * 3),
	}

	lights := []bool{true, false, true}

	fmt.Printf("%s\n", pizza)
	fmt.Printf("%d\n", graduations)
	fmt.Printf("%v\n", departures)
	fmt.Printf("%t\n", lights)

}
