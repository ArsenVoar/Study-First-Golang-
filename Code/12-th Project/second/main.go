package main

import (
	"fmt"
	"time"
)

func main() {
	nilChanel()

	unbufferedChannel()

	bufferedChannel()

	forRange()
}

func forRange() {
	bufferedChannel := make(chan int, 3)

	numbers := []int{5, 6, 7, 8}

	//show all elements with for

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for {
		v, ok := <-bufferedChannel
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}

	bufferedChannel = make(chan int, 3)

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for v := range bufferedChannel {
		fmt.Println("buffered", v)
	}
}

func bufferedChannel() {
	bufferedChannel := make(chan int, 2)
	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	bufferedChannel <- 1
	bufferedChannel <- 2

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	bufferedChannel <- 3

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChannel), cap(bufferedChannel))

	fmt.Println(<-bufferedChannel)
}

func unbufferedChannel() {

	unnbufferedChannel := make(chan int)
	fmt.Printf("Len: %d Cap: %d\n", len(unnbufferedChannel), cap(unnbufferedChannel))

	//blocks on reading then write
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)

		unnbufferedChannel <- 1
	}(unnbufferedChannel)

	value := <-unnbufferedChannel
	fmt.Println(value)

	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)

		value := <-chanForReading

		fmt.Println(value)
	}(unnbufferedChannel)

	unnbufferedChannel <- 2
}

func nilChanel() {
	var nilChannel chan int
	fmt.Printf("Len: %d Cap:%d\n", len(nilChannel), cap(nilChannel))
}
