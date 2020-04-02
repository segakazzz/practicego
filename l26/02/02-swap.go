package main

import "fmt"

func main() {

	a, b := 1.23, 2.34
	// swap(&a, &b)
	fmt.Printf("a: %f b: %f\n", a, b)

	var c, d *float64
	c, d = &a, &b
	fmt.Printf("c: %f, %p d: %f, %p\n", *c, c, *d, d)
	c, d = swap2(c, d)
	fmt.Printf("c: %f, %p d: %f, %p\n", *c, c, *d, d)

}

func swap(pa *float64, pb *float64) {
	*pa, *pb = *pb, *pa
}

func swap2(c, d *float64) (*float64, *float64) {
	return d, c
}
