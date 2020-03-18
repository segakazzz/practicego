package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		m   float64
		err error
	)
	if len(os.Args) < 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	if m, err = strconv.ParseFloat(os.Args[1], 64); err != nil {
		fmt.Printf("I coudn't get %q, sorry\n", os.Args[1])
		return
	}

	switch {
	case m < 2.0:
		fmt.Printf("%.2f is micro \n", m)
	case m < 3.0:
		fmt.Printf("%.2f is very minor \n", m)
	case m < 4.0:
		fmt.Printf("%.2f is minor \n", m)
	case m < 5.0:
		fmt.Printf("%.2f is light \n", m)
	case m < 6.0:
		fmt.Printf("%.2f is moderate \n", m)
	case m < 7.0:
		fmt.Printf("%.2f is strong \n", m)
	case m < 8.0:
		fmt.Printf("%.2f is major \n", m)
	case m < 10.0:
		fmt.Printf("%.2f is great \n", m)
	default:
		fmt.Printf("%.2f is massive \n", m)

	}

}
