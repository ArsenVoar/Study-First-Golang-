package main

import (
	"fmt"
)

//var name string //- var scope example

//const name = "user" - constant

func main() {
	name := "arsen"
	hello := fmt.Sprintf("hello %s", name)
	fmt.Println(hello)

	fmt.Println(name) // - var scope example

	/*	ourBool := true // - short syntax
		var hello string = "hello"// - (string) without "= "hello"" will be empty string */
	var ourBool bool // - (bool)

	fmt.Println(name)    // - constant
	fmt.Println(hello)   // - (string)*/
	fmt.Println(ourBool) // - (bool)

	hello = "hello"

	fmt.Printf("Type: %T Value:%v\n", name, name)
	fmt.Printf("Type: %T Value:%v\n", name, name)       // - constant
	fmt.Printf("Type: %T Value:%v\n", name, name)       // - var scope example
	fmt.Printf("Type: %T Value:%v\n", hello, hello)     // - (string)
	fmt.Printf("Type: %T Value:%v\n", ourBool, ourBool) // -/bool
}
