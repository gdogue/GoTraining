package main

import (
	"fmt"
)

func main() {

	s := struct {
		name    string
		friends map[string]int
		drinks  []string
	}{
		name: "James",
		friends: map[string]int{
			"Money": 111,
			"Q":     222,
			"M":     333,
		},
		drinks: []string{"water", "soda", "coffee"},
	}

	fmt.Println(s)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for _, v := range s.drinks {
		fmt.Println(v)
	}
}
