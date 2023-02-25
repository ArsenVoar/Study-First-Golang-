package main

import "fmt"

func main() {
	definition()
}

type SNOW struct {
	Number int
}

func definition() {
	square := SNOW{Number: 4}
	pSquare := &square

	square2 := SNOW{Number: 2}

	square.Answer()
	square2.Answer()

	pSquare.Answer()
	square.Answer()
	pSquare.Answer()

	square.Answer()
}

func (p SNOW) Answer() {
	fmt.Printf("%T, %#v \n", p, p)
	fmt.Printf("SNOW Answer: %d \n", p.Number*4)
} // - Value reciever

func (p *SNOW) Ans(number2 int) {
	fmt.Printf("%T, %#v \n", p, p)
	p.Number *= number2
	fmt.Printf("%T, %#v \n", p, p)
} // - Pointer reciever
