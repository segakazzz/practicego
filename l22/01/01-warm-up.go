package main

import "fmt"

func main() {
	map1 := map[string]string{
		"a": "123456789",
		"b": "4325452323",
	}
	map2 := map[int]string{
		132: "123456789",
		454: "4325452323",
	}
	map3 := map[string]bool{
		"a": true,
		"b": false,
	}
	map4 := map[int]map[int]int{
		453: map[int]int{
			232: 343,
			23:  354,
		},
		2334: map[int]int{
			22:  43,
			243: 2454,
		},
	}

	fmt.Printf("%#v\n", map1)
	fmt.Printf("%#v\n", map2)
	fmt.Printf("%#v\n", map3)
	fmt.Printf("%#v\n", map4)

}
