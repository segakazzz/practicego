package main

import (
	"encoding/json"
	"fmt"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type gameEncodable struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Genre string `json:"genre"`
}

func newGame(id, price int, name, genre string) game {
	return game{item: item{id: id, name: name, price: price}, genre: genre}
}

func addGame(games []game, game game) []game {
	return append(games, game)
}

func load() []game {
	return append([]game{},
		newGame(1, 50, "god of war", "action adventure"),
		newGame(2, 40, "x-com 2", "strategy"),
		newGame(5, 20, "minecraft", "sandbox"),
	)
}

func loadFromString(str string) []game {
	gamesEnc := []gameEncodable{}
	json.Unmarshal([]byte(str), &gamesEnc)
	fmt.Println(gamesEnc)

	games := []game{}
	for _, g := range gamesEnc {
		games = append(games, game{item: item{id: g.ID, name: g.Name, price: g.Price}, genre: g.Genre})
	}

	return games
}

func indexByID(games []game) map[int]game {
	newMap := map[int]game{}
	for _, g := range games {
		newMap[g.id] = g
	}
	return newMap
}

func printGame(game game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		game.id, game.name, "("+game.genre+")", game.price)
}
