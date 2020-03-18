package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		start int64
		end   int64
		err   error
		sum   int
		str   string
	)

	if len(os.Args) < 3 {
		fmt.Println("Usage: command start_num end_num")
		return
	}
	if start, err = strconv.ParseInt(os.Args[1], 10, 64); err != nil {
		fmt.Printf("Invalid start number entered: %q\n", start)
		return
	}

	if end, err = strconv.ParseInt(os.Args[2], 10, 64); err != nil {
		fmt.Printf("Invalid end number entered: %q\n", end)
		return
	}

	for i := int(start); i <= int(end); i++ {
		if i%2 == 1 {
			continue
		}
		sum += i
		str += strconv.Itoa(i)
		if i != int(end) {
			str += " + "
		} else {
			str += " = "
		}
	}
	str += strconv.Itoa(sum)
	fmt.Println(str)
}
