package main

import (
	"fmt"
	"runtime" // --- 1 ---
)

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)

	// |||||||||||||||||||||||||
	fmt.Println(some(2, 3))

	// deferValues()
	makePanic()

	go showNumber(22) // --- 1 ---

	// ||||||||||||||||||||||||||||||

	runtime.GOMAXPROCS(6)
	fmt.Println(runtime.NumCPU())

	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

	//runtime.Gosched() - try it
	//time.Sleep(time.Second) - try it too --- 2 ---

	fmt.Println("exit")
}

func makePanic() {
	defer func() {
		panicvalue := recover()
		fmt.Println(panicvalue)
	}()

	panic("some panic")
}

func deferValues() {
	for j := 0; j < 5; j++ {
		defer fmt.Println("first", j)
	}
	for j := 0; j < 5; j++ {
		defer func() {
			fmt.Println("second", j)
		}()
	}
}

// ||||||||||||||||||||||||||||||||||||||||||
func some(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()
	sum = x + y
	return
}

// ||||||||||||||||||||||||||||||||||||||||

func showNumber(num int) {
	for p := 0; p < num; p++ {
		fmt.Println(p)
	}
}
