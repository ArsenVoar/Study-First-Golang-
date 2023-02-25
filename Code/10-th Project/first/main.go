package main

import "fmt"

func main() {
	getSlice()
	copySlice()
}

func copySlice() {
	destination := make([]string, 0, 2)
	source := []string{"Paul", "Mike"}

	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("Type: %T Value: %#v\n", destination, destination)
	fmt.Printf("Length: %d Capacity: %d\n", len(destination), cap(destination))
}

func getSlice() {
	intArr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	intSlice := [...]int{1: 3}
	fmt.Printf("Type: %T Value: %#v\n", intSlice, intSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(intSlice), cap(intSlice))

	fullSlice := intArr[:] // IntArr[0:5]
	fmt.Printf("Type: %T Value: %#v\n", fullSlice, fullSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(fullSlice), cap(fullSlice))

	sliceFromSlice := fullSlice[:3]
	fmt.Printf("Type: %T Value: %#v\n", sliceFromSlice, sliceFromSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(sliceFromSlice), cap(sliceFromSlice))

	intArr[2] = 88 // if we change
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)
	fmt.Printf("Type: %T Value: %#v\n", intSlice, intSlice)
	fmt.Printf("Type: %T Value: %#v\n", fullSlice, fullSlice)
	fmt.Printf("Type: %T Value: %#v\n", sliceFromSlice, sliceFromSlice)
}
