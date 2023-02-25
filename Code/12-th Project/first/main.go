package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	withoutWait()

	withWait()

	wrongAdd()

	withoutCuncurent()

	withoutMutex()

	withMutex()

	readWithMutex()

	readWithRWMutex()
}

func readWithRWMutex() {
	start := time.Now()
	var (
		counter int
		WG      sync.WaitGroup
		mu      sync.RWMutex
	)

	WG.Add(100)

	for t := 0; t < 50; t++ {
		go func() {
			defer WG.Done()

			mu.RLock()

			time.Sleep(time.Nanosecond)
			_ = counter

			mu.RUnlock()
		}()
	}

	go func() {
		defer WG.Done()

		mu.Lock()

		time.Sleep(time.Nanosecond)
		counter++

		mu.Unlock()
	}()
	WG.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		WG      sync.WaitGroup
		mu      sync.Mutex
	)

	WG.Add(100)

	for t := 0; t < 50; t++ {
		go func() {
			defer WG.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			_ = counter

			mu.Unlock()
		}()
	}

	go func() {
		defer WG.Done()

		mu.Lock()

		time.Sleep(time.Nanosecond)
		counter++

		mu.Unlock()
	}()
	WG.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func withMutex() {
	start := time.Now()
	var counter int
	var WG sync.WaitGroup
	var Mu sync.Mutex

	WG.Add(1000)

	for u := 0; u < 1000; u++ {
		go func() {
			defer WG.Done()
			time.Sleep(time.Nanosecond)

			Mu.Lock()
			counter++
			Mu.Unlock()
		}()
	}
	WG.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func withoutMutex() {
	start := time.Now()
	var counter int
	var WG sync.WaitGroup

	WG.Add(1000)

	for u := 0; u < 1000; u++ {
		go func() {
			defer WG.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()
	}
	WG.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func withoutCuncurent() {
	start := time.Now()
	var counter int

	for u := 0; u < 1000; u++ {
		time.Sleep(time.Nanosecond)
		counter++
	}
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func wrongAdd() {
	var WG sync.WaitGroup

	for h := 0; h < 8; h++ {

		go func(h int) {
			//defer wg.Done()
			WG.Add(1)

			fmt.Println(h + h)
			WG.Done()
		}(h)
	}
	WG.Wait()
	fmt.Println("exit")
}

func withWait() {
	var WG sync.WaitGroup

	for h := 0; h < 8; h++ {
		WG.Add(1)

		go func(h int) {
			//defer wg.Done()

			fmt.Println(h + h)
			WG.Done()
		}(h)
	}
	WG.Wait()
	fmt.Println("exit")
}

func withoutWait() {
	for k := 0; k < 8; k++ {
		go fmt.Println(k + k)
	}
	fmt.Println("exit")
}
