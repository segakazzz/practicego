package main

import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	type (
		Feet float64
		Meters float64
	)

	var (
		feet Feet
		meters Meters
	)

	feet_float64, _ := strconv.ParseFloat(os.Args[1], 64)
	feet = Feet(feet_float64)
	meters = Meters(feet) * 0.3048
	fmt.Println("%g feet is %g meters.\n", feet, meters)

}