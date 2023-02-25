package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	baseSelect()
	graqcefulShutDown()
}

func graqcefulShutDown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	for {
		select {
		case <-timer:
			fmt.Println("time up")
			return

		case sig := <-sigChan:
			fmt.Println("Stopped by signal", sig)
			return // in cmd write "go run main.go" and push ctrl + c, this will be a signal
		}
	}
}

func baseSelect() {
	bufferedChan := make(chan string, 3)
	bufferedChan <- "first"

	select {
	case point := <-bufferedChan:
		fmt.Println("read", point)

	case bufferedChan <- "second":
		fmt.Println("write", bufferedChan, bufferedChan)
	}

	unbufChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		unbufChan <- 1
	}()
	select {

	//	case bufferedChan <- "third":
	//		fmt.Println("unblocking writing")

	case val := <-unbufChan:
		fmt.Println("blocking reading", val)
	case <-time.After(time.Millisecond * 500):
		fmt.Println("time is up")

		//		default:
		//			fmt.Println("default case")
	}
}
