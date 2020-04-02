package main

import (
	"fmt"
	"strconv"
	"strings"
)

func runCmd(input string, games []game, gameMap map[int]game) bool {

	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return false
	}

	switch cmd[0] {
	case "quit":
		return cmdQuit()
	case "list":
		return cmdList(games)
	case "id":
		return cmdByID(cmd, games, gameMap)
	}
	return true
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
		return false
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return false
	}

	g, ok := gameMap[id]
	if !ok {
		fmt.Println("sorry. i don't have the game")
		return false
	}

	printGame(g)

	return true
}
