package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		sum int
		num int
		err error
		cnt int
	)

	args := os.Args[1:]
	if len(args) == 0 || len(args) > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers)")
		return
	}

	fmt.Printf("Your numbers: %v\n", args)
	for _, v := range args {
		num, err = strconv.Atoi(v)
		if err == nil {
			sum += num
			cnt++
		}
	}
	fmt.Printf("Averge %.2f\n", float64(sum)/float64(cnt))

}
