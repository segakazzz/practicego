package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(strings.ToLower(name))
}
