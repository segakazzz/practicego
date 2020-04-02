package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var lineNumber = 0
	in := bufio.NewScanner(os.Stdin)
	domains := map[string]int{}
	totalVisit := 0

	for in.Scan() {
		lineNumber++
		line := in.Text()
		lineSlice := strings.Fields(line)
		if len(lineSlice) != 2 {
			fmt.Printf("wring input [%s] (line #%d)\n", line, lineNumber)
			return
		}

		visits, err := strconv.Atoi(lineSlice[1])
		if err != nil {
			fmt.Printf("wring input %q (line #%d)\n", lineSlice[1], lineNumber)
			return
		}

		if visits < 0 {
			fmt.Printf("wring input %q (line #%d)\n", lineSlice[1], lineNumber)
			return
		}

		_, ok := domains[lineSlice[0]]
		if ok {
			domains[lineSlice[0]] += visits
		} else {
			domains[lineSlice[0]] = visits
		}
	}

	fmt.Printf("%30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("=", 50))

	for k, v := range domains {
		totalVisit += v
		fmt.Printf("%30s %10d\n", k, v)
	}

	fmt.Println()
	fmt.Printf("%30s %10d\n", "TOTAL", totalVisit)
}
