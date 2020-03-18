package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	val, _ := strconv.ParseInt(os.Args[1], 10, 8)
	fmt.Printf("%T\n", val)
	fmt.Println("int8 value is:", int8(val))
	val1, _ := strconv.ParseInt(os.Args[1], 10, 16)
	fmt.Printf("%T\n", val1)
	fmt.Println("int16 value is:", int16(val1))
	val2, _ := strconv.ParseInt(os.Args[1], 10, 32)
	fmt.Printf("%T\n", val2)
	fmt.Println("int32 value is:", int32(val2))
	val3, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Printf("%T\n", val3)
	fmt.Println("int64 value is:", int64(val3))

	val4, _ := strconv.ParseInt(os.Args[1], 2, 8)
	fmt.Println("int8 value is:", int8(val4))
}
