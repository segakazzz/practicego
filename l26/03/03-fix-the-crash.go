package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	var c *computer = &computer{}
	change(c, "apple")
	fmt.Printf("brand: %s\n", c.brand)
}

func change(c *computer, brand string) {
	c.brand = &brand
}
