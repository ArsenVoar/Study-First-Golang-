package main

import (
	"fmt"
)

func main() {

	first, second := 1, 2

	var person func(x, y int) int

	person = func(x, y int) int { return x * y }

	fmt.Println(person(first, second))

	player := func(x, y int) int { return x / y } // - short var witn func
	fmt.Println(player(second, first))            //   - how to use var with func
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

func createdevider(divider int) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}
