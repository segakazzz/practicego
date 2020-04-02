package main

import "fmt"

func main() {

	phoneNumbers := make(map[string]string)
	phoneNumbers["bowen"] = "202-555-0179"
	phoneNumbers["dulin"] = "03.37.77.63.06"
	phoneNumbers["greco"] = "03489940240"

	fmt.Printf("%s\n", phoneNumbers["bowen"])

	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	fmt.Printf("%t\n", products[879401371])

	phoneNumbers2 := map[string][]string{
		"bowen": []string{"202-555-0179"},
		"dulin": []string{"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": []string{"03489940240", "03489900120"},
	}

	fmt.Printf("%#v\n", phoneNumbers2["greco"])
}
