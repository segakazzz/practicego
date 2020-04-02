package main

import (
	"encoding/json"
	"fmt"
)

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type jsonEncodable struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Genre string `json:"genre"`
}

func main() {

	jsonGames := []jsonEncodable{}
	err := json.Unmarshal([]byte(data), &jsonGames)
	if err != nil {
		fmt.Println(err)
		return
	}

	games := []game{}
	for _, g := range jsonGames {
		games = append(games, game{item: item{id: g.ID, name: g.Name, price: g.Price}, genre: g.Genre})
	}
	fmt.Printf("%#v\n", games)
}
