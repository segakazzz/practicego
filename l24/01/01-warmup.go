package main

import "fmt"

func main() {

	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	games := []game{
		{
			item: item{
				id:    1,
				name:  "god of war",
				price: 50,
			},
			genre: "action adventure",
		},
		{
			item: item{
				id:    2,
				name:  "x-com 2",
				price: 30,
			},
			genre: "strategy",
		},
		{
			item: item{
				id:    3,
				name:  "minecraft",
				price: 20,
			},
			genre: "sandbox",
		},
	}

	fmt.Printf("%-10s %-20s %-10s %-20s\n", "id", "name", "price", "genre")
	for _, game := range games {
		fmt.Printf("%-10d %-20s %-10d %-20s\n", game.id, game.name, game.price, game.genre)
	}

	// fmt.Printf("%#v\n", games)
}
