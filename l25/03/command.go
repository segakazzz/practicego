package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func runCmd(input string, games []game, gameMap map[int]game) bool {

	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return true
	}
	switch cmd[0] {
	case "quit":
		return cmdQuit()
	case "list":
		return cmdList(games)
	case "id":
		return cmdByID(cmd, games, gameMap)
	case "save":
		return cmdSave(games)
	default:
		return true
	}
}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

func cmdByID(cmd []string, games []game, gameMap map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := gameMap[id]
	if !ok {
		fmt.Println("sorry. i don't have the game")
		return true
	}

	printGame(g)

	return true
}

func cmdSave(games []game) bool {
	gamesEncodable := []gameEncodable{}

	for _, g := range games {
		gamesEncodable = append(gamesEncodable, gameEncodable{
			ID: g.id, Name: g.name, Price: g.price, Genre: g.genre,
		})
	}
	js, _ := json.MarshalIndent(gamesEncodable, "", "\t")
	fmt.Println(string(js))
	return true
}
