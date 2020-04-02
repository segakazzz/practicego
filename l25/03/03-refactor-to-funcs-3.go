package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// games := load()

	games := loadFromString(`
[
        {
                "id": 1,
                "name": "god of war",
                "price": 50,
                "genre": "action adventure"
        },
        {
                "id": 2,
                "name": "x-com 2",
                "price": 40,
                "genre": "strategy"
        },
        {
                "id": 5,
                "name": "minecraft",
                "price": 20,
                "genre": "sandbox"
        }
]`)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > quit   : quits
  > save   : print json
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		if !runCmd(in.Text(), games, indexByID(games)) {
			break
		}
	}
}
