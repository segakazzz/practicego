package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	val, _ := strconv.ParseInt(os.Args[1], 10, 64)
	t, _ := time.ParseDuration("1h30m")

	fmt.Println(t * time.Duration(val))
}
