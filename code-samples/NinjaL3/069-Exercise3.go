package main

import (
	"fmt"
)

func main() {

	yr := 1964

	for yr >= 1964 && yr <= 2021 {
		fmt.Println(yr)
		yr++
	}
}
