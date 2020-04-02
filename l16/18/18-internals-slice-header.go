package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	debug.SetGCPercent(-1)
	report("initial memory usage")

	elem := [size]int{}
	report("memory usage after 10m int elements")
	copyElem := elem
	report("after copying 10m int elements")
	passArray(copyElem)

	sliceElem := elem[:]
	report("after copying to slice")

	sliceElem2 := elem[1000:10000]
	report("after copying to slice2")

	fmt.Println(unsafe.Sizeof(elem))
	fmt.Println(unsafe.Sizeof(copyElem))
	fmt.Println(unsafe.Sizeof(sliceElem))
	fmt.Println(unsafe.Sizeof(sliceElem2))

	_, _, _, _ = elem, copyElem, sliceElem, sliceElem2
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
