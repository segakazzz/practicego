package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	sHeader := "Location,Size,Beds,Baths,Price"
	sData := `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

	separator := ","

	header := strings.Split(sHeader, separator)

	locations := []string{}
	sizes := []int{}
	beds := []int{}
	baths := []int{}
	prices := []int{}

	var number int
	var err error

	rows := strings.Split(sData, "\n")
	for _, r := range rows {
		rowData := strings.Split(r, separator)
		for i, v := range rowData {
			number, err = strconv.Atoi(v)
			if i != 0 && err == nil {
				switch i {
				case 1:
					sizes = append(sizes, number)
				case 2:
					beds = append(beds, number)
				case 3:
					baths = append(baths, number)
				case 4:
					prices = append(prices, number)
				}
			}
			if i == 0 {
				locations = append(locations, v)
			}
		}
	}

	for _, v := range header {
		fmt.Printf("%10s", v)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 50))

	for i := range sizes {
		fmt.Printf("%10s", locations[i])
		fmt.Printf("%10d", sizes[i])
		fmt.Printf("%10d", beds[i])
		fmt.Printf("%10d", baths[i])
		fmt.Printf("%10d", prices[i])
		fmt.Println()
	}

	fmt.Println(strings.Repeat("=", 50))

	fmt.Printf("%10s", "")
	fmt.Printf("%10.2f", Average(sizes))
	fmt.Printf("%10.2f", Average(beds))
	fmt.Printf("%10.2f", Average(baths))
	fmt.Printf("%10.2f", Average(prices))
	fmt.Println()

}

// Average ...
// Return Average of int
func Average(slice []int) float64 {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return float64(sum) / float64(len(slice))
}
