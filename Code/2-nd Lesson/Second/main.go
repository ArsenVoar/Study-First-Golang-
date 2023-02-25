package main

import (
	"fmt"
)

func main() {
	Greet()
	//	funcName()
	simple()
}

/*func funcName(Settings) {
} //- struct of function
*/
// |||||||||||||||||||||||||||||||||||
func Greet() {
	fmt.Println("Hello")
} // -simple function

// ||||||||||||||||||||||||||||||
func simple() {
	var name string

	fmt.Printf("Hello %v\n", name)
}
