package main

import "fmt"

func main() {
	variadic()
	SliceStructure()
	passSlicetoFunction()
	sliceWithNew()
}

func sliceWithNew() {
	slicePointer := new([]int)

	fmt.Printf("Type: %T Value: %#v\n", slicePointer, slicePointer)
	fmt.Printf("Length: %d Capacity: %d\n", len(*slicePointer), cap(*slicePointer))

	newSlice2 := append(*slicePointer, 1)
	fmt.Printf("Type: %T Value: %#v\n", newSlice2, newSlice2)
	fmt.Printf("Length: %d Capacity: %d\n", len(newSlice2), cap(newSlice2)) // it's now good, but we can (make better)
}

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

func passSlicetoFunction() {
	initialSlice := []int{1, 2}
	fmt.Printf("Type: %T Value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(initialSlice), cap(initialSlice))

	changeValue(initialSlice)

	fmt.Printf("Type: %T Value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(initialSlice), cap(initialSlice))

	newSlice := append(initialSlice, 3)
	fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(newSlice), cap(newSlice)) // capacity will x2

	newSlice = appendValue(newSlice)
	fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(newSlice), cap(newSlice))

	fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(newSlice), cap(newSlice)) // --- or this --- 2 ---

	//	newSlice2 := appendValue(newSlice)
	//	fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice)
	//	fmt.Printf("Length: %d Capacity: %d\n", len(newSlice), cap(newSlice))

	//	fmt.Printf("Type: %T Value: %#v\n", newSlice2, newSlice2)
	//	fmt.Printf("Length: %d Capacity: %d\n", len(newSlice2), cap(newSlice2)) // --- we can do this --- 1 ---, but better 2, because in this way newSlice if we will work with hit willn't have changes
}

func changeValue(slice []int) {
	slice[1] = 15
}

func appendValue(slice []int) []int {
	slice = append(slice, 4, 5)
	fmt.Printf("Type: %T Value: %#v\n", slice, slice)
	fmt.Printf("Length: %d Capacity: %d\n", len(slice), cap(slice))
	return slice
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

func SliceStructure() {
	initialSlice := []int{7, 8}
	fmt.Printf("Type: %T Value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(initialSlice), cap(initialSlice))

	initialArray := (*[2]int)(initialSlice)
	fmt.Printf("Type: %T Value: %#v\n", initialArray, initialArray)
	fmt.Printf("Length: %d Capacity: %d\n", len(initialArray), cap(initialArray))
}

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

func variadic() {
	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 4, 5, 6, 7, 8)

	firstSlice := []int{5, 6, 7, 8}
	secondSlice := []int{9, 3, 2, 1}
	showAllElements(firstSlice...)

	newSlice := append(firstSlice, secondSlice...)
	fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice) // (firstSlice, 9, 3, 2, 1)
}

func showAllElements(values ...int) {
	for _, val := range values {
		fmt.Println("value", val)
	}
	fmt.Println()
}
