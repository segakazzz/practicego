package main

import "fmt"

func main() {
	const word = "console"

	for _, s := range word {
		fmt.Printf("decimal: %[1]d hexadecimals: %[1]x\n", s)
	}

	bword := []byte(word)
	bword[1] = 'x'
	fmt.Println(string(bword))

	bword[2] = 99
	fmt.Println(string(bword))

	bword[3] = 0x65
	fmt.Println(string(bword))

}
