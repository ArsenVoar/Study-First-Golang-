package main

import (
	"fmt"
)

type Our string

func (t Our) Hello() {
	fmt.Println("Hello")
}

func main() {
	definition()
}

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("Figure Perimeter: %d \n", s.Side*4)
}

func (s Square) WrongScale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func definition() {
	square := &Square{Side: 4}

	square2 := Square{Side: 2}

	square.Perimeter()
	square2.Perimeter()
}
