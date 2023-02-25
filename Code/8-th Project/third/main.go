package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// arrays()
	slices()
}

//slices
func slices() {
	var defaultSlices []int
	fmt.Printf("Type: %T Value: %#v\n", defaultSlices, defaultSlices)

	//defaultSlices[0] = 1

	slice := make([]int, 0, 5)
	fmt.Printf("Type: %T Value: %#v\n", slice, slice)
	fmt.Printf("Length: %d Capacity: %d\n", len(slice), cap(slice))
}

func arrays() {
	//arrays
	var intArr [3]int
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	// index
	intArr[0] = 4
	intArr[1] = 8
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)

	//syntax cuts
	people := [2]Person{
		{
			Name: "Tate",
			Age:  25,
		},
		{
			Age:  27,
			Name: "Mack",
		},
	}
	fmt.Printf("Type: %T Value: %#v\n", people, people)

	//automatic calculation
	stringArr := [...]string{"One", "Two", "Eight"}
	fmt.Printf("Type: %T Value: %#v\n", stringArr, stringArr)

	fmt.Printf("Length: %d Capacity: %d/n", len(stringArr), cap(stringArr))

	for inx, value := range stringArr {
		fmt.Printf("Index: %d Value: %s\n", inx, value)
	}
}
