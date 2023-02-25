package main

import "fmt"

func main() {

	//empty value
	var turn1 *int
	fmt.Println(justTurn(turn1))

	turn2 := 0
	fmt.Println(justTurn(&turn2))

	turn3 := 10
	fmt.Println(justTurn(&turn3))

	// pointers usage
	//side effect

	num := 4
	//1...square(num)
	fmt.Println(num)

	//2...squarePointer(&num)
	fmt.Println(num)
}
func justTurn(door *int) bool {
	return door != nil
}

// ||||||||||||||||||||||||||||||||||||||||||||||

//1...
func square(num int) {
	num *= num
	fmt.Println(num)
}

// ||||||||||||||||||||||||||||||||||||||||||

//2...
func squarePointer(num *int) {
	*num *= *num
}
