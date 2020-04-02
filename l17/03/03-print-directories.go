package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	var buffer []byte
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	for _, d := range args {
		files, err := ioutil.ReadDir(d)
		if err != nil {
			fmt.Println(err)
			return
		}
		buffer = append(buffer, []byte(d+"\n")...)
		for _, file := range files {
			if file.IsDir() {
				buffer = append(buffer, []byte("\tDirectory...")...)
			} else {
				buffer = append(buffer, []byte("\tFile...")...)
			}
			buffer = append(buffer, []byte(file.Name()+"\n")...)
		}
	}

	err := ioutil.WriteFile("dirs.txt", buffer, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
