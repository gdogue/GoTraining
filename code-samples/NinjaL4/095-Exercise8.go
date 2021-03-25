package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{}

	m["James"] = []string{"Bond", "007", "Agent"}
	m["Money"] = []string{"Penny", "Miss", "Secretary"}
	m["Dr No"] = []string{"Bad", "Guy", "Evil"}

	//fmt.Println(m)

	for k, v := range m {
		fmt.Println("Key:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
