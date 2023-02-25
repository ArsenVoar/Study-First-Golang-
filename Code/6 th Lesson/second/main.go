package main

import "fmt"

type Runner interface {
	Run() string
} // - interface

func main() {
	interfaceValue()
}

type WorkOut interface {
	WorkOut() string
}

type PowerLifter interface {
	PowerLifter() string
}

type Sport interface {
	WorkOut
	PowerLifter
}

type Sportsmen interface {
	Name() string
}

// WorkOut = int64(1) // cannot working

func interfaceValue() {
	var WorkOut WorkOut

	fmt.Printf("Type: %T Value: %#v\n", WorkOut, WorkOut)

	if WorkOut == nil {
		fmt.Println("Nil")
	}

	var unnamed *Sportsmen
	fmt.Printf("Type: %T Value: %#v\n", unnamed, unnamed)

	fmt.Printf("Type: %T Value: %#v\n", WorkOut, WorkOut)
	if WorkOut == nil {
		fmt.Println("WorkOut is Nil")
	} else {
		fmt.Println("No Nil")
	}
}
