package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var n int
	var err error
	args := os.Args[1:]
	for _, s := range args {
		n, err = strconv.Atoi(s)
		if err == nil && isPrime(n) {
			fmt.Printf("%5d ", n)
		}
	}
	fmt.Println()

	type (
		newType int
	)

	var value newType = 1
	var value2 int = 1
	fmt.Println(int(value) == value2)

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
