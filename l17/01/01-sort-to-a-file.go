package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {

	var slice []string
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	for _, v := range args {
		slice = append(slice, v)
	}
	sort.Strings(slice)

	fmt.Println(slice)

	var buffer []byte
	for _, v := range slice {
		buffer = append(buffer, []byte(v)...)
		buffer = append(buffer, []byte("\n")...)
	}
	err := ioutil.WriteFile("sorted.txt", buffer, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}
