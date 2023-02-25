package main

import "fmt"

func main() {

	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer) // - default *

	// variable
	var any int32 = 7
	fmt.Printf("%T %#v \n", any, any)

	var pointerAny *int32 = &any
	fmt.Printf("%T %#v %#v \n", pointerAny, pointerAny, *pointerAny)

	//get pointer via new keyword
	var newPointer = new(float32)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 3
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
}
