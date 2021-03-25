package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{
		"James": {"Bond", "007", "Agent"},
		"Money": {"Penny", "Miss", "Secretary"},
		"Dr No": {"Bad", "Guy", "Evil"},
	}

	m["Gary"] = []string{"Rich", "001", "Student"}

	delete(m, "James")

	for k, v := range m {
		fmt.Println("Key:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
