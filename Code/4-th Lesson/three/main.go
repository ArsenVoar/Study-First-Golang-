package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	max = 6
	min = 1
)

func main() {

	rand.Seed(time.Now().UnixNano())

	value := rand.Intn(max-min) + min

	switch {
	case value > 2:
		fmt.Printf("Valus %d greater than 2 \n", value)
	case value < 2:
		fmt.Printf("Valus %d less than 2 \n", value)
	default:
		fmt.Println("value equals 2")
	}

	getFour := 4
	switch number := rand.Intn(max-min) + min; number {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")

	case getFour:
		fmt.Println("Four")
		fallthrough
	default:
		fmt.Println("Default case")
	}

	switch value {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")

	case getFour:
		fmt.Println("Four")
	default:
		fmt.Println("Default case")
	} //  - switch
}
