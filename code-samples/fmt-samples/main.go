package main

import (
	"fmt"
)

var y = 42

func main() {

	//Format value to print to stdout
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	//Using multiple args
	y = 911
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	//Printing to a string and then using that string
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	//Print the value as is
	fmt.Printf("%v\n", y)
}
