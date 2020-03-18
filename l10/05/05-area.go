package main

import "fmt"

func main(){
	const (
		width int16 = 25
		height int32 = int32(width) * 2
	)
	fmt.Printf("area = %d\n", int32(width) * height)
}