package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum, str := 0, ""
	for i := 1; i <= 10; i++ {
		sum += i
		str += strconv.Itoa(i)
		if i != 10 {
			str += " + "
		} else {
			str += " = "
		}
	}
	str += strconv.Itoa(sum)
	fmt.Println(str)
}
