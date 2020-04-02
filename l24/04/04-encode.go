package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

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
		Genre string `json:"gemre"`
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

quit:
	for in.Scan() {
		text := in.Text()
		input := strings.Fields(text)
		if len(input) > 0 {
			switch input[0] {
			case "quit":
				fmt.Println("quit command entered")
				break quit
			case "list":
				fmt.Printf("%-10s %-20s %-10s %-20s\n", "id", "name", "price", "genre")
				for _, game := range games {
					fmt.Printf("%-10d %-20s %-10d %-20s\n", game.id, game.name, game.price, game.genre)
				}
				fmt.Println("list command entered")

			case "id":
				if len(input) == 1 {
					fmt.Println("wrong id entered")
				} else if id, err := strconv.Atoi(input[1]); err != nil {
					fmt.Printf("wrong id entered %q\n", input[1])
				} else {
					for _, game := range games {
						if game.id == id {
							fmt.Printf("%-10d %-20s %-10d %-20s\n", game.id, game.name, game.price, game.genre)
							break
						}
					}
					fmt.Printf("sorry I don't have a game id:%d \n", id)
				}

			case "save":
				jsonGames := []jsonEncodable{}
				for _, game := range games {
					jsonGames = append(jsonGames, jsonEncodable{ID: game.id, Name: game.name, Price: game.price, Genre: game.genre})
				}
				js, err := json.MarshalIndent(jsonGames, "", "\t")
				fmt.Printf("%s\n", js)
				if err != nil {
					fmt.Println(err)
					return
				}

			default:
				fmt.Printf("%s\n", text)
			}

		}
	}
}
