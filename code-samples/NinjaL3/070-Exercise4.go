package main

import (
	"fmt"
)

func main() {

	yr := 1964

	for {
		if yr > 2021 {
			break
		}
		fmt.Println(yr)
		yr++
	}

}
