package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Requires age")
	} else {
		age, _ := strconv.ParseInt(os.Args[1], 10, 64)
		if age < 0 {
			fmt.Printf("Wrong age: %q\n", age)
		} else if age > 17 {
			fmt.Println("R-Raged")
		} else if age >= 13 && age <= 17 {
			fmt.Println("PG-13")
		} else {
			fmt.Println("PG-Rated")
		}
	}

}
