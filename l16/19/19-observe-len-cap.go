package main

import "fmt"

func main() {
	// games := []string{
	// 	"mario", "kinopoko", "kuppa",
	// }

	// games := []string{}
	// games = append(games, "pacman", "mario", "tetris", "doom")
	// games = append(games, "one more")

	// fmt.Printf("len: %d cap: %d games: %q\n", len(games), cap(games), games)

	games := []string{"pacman", "mario", "tetris", "doom"}

	for i := range games {
		fmt.Printf("len: %d cap: %d games: %q\n", len(games[i:i+1]), cap(games[i:i+1]), games[i:i+1])
	}

	zero := games[:0]
	for i := 0; i < len(games)+1; i++ {
		fmt.Println(i)
		zero = append(zero, "kuppa")
		fmt.Printf("len: %d cap: %d games: %q\n", len(zero), cap(zero), zero)
		fmt.Printf("len: %d cap: %d games: %q\n", len(games), cap(games), games)
	}

}
