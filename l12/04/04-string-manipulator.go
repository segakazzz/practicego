package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println(`
[command] [string]
		
Available commands: lower, upper and title
`)
		return
	}

	t, s := os.Args[1], os.Args[2]

	switch t {
	case "lower":
		fmt.Printf("%s\n", strings.ToLower(s))
	case "upper":
		fmt.Printf("%s\n", strings.ToUpper(s))
	case "title":
		fmt.Printf("%s\n", strings.Title(s))
	default:
		fmt.Printf("Unknown command: %q\n", t)
	}
}
