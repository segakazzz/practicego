package main

import "fmt"

func main() {
	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json := `
{
	"Items":[
		"Item":{
			"name": "Teddy Bear"
		}
	]
}
`
	fmt.Println(json)
}
