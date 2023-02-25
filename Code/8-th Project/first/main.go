package main

import "fmt"

func main() {
	interfacesValues()
}

type Runner interface {
	Run() string
} // interface

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Dycker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func interfacesValues() {
	var runner Runner
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("runner is nil")
	}
	// runner = int64(1) // cannot work
	//runner.Run() // too /8/

	var unnamedRunner *Human
	fmt.Printf("Type: %T Value: %#v\n", unnamedRunner, unnamedRunner)

	namedRunner := &Human{Name: "Chris"}
	fmt.Printf("Type: %T Value: %#v\n", namedRunner, namedRunner)

	var empty interface{} = unnamedRunner
	fmt.Printf("Type: %T Value: %#v\n", empty, empty)

	empty = int64(1)
	fmt.Printf("Type: %T Value: %#v\n", empty, empty)

	empty = true
	fmt.Printf("Type: %T Value: %#v\n", empty, empty)
}
