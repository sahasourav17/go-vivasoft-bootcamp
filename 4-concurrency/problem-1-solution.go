package main

import (
	"fmt"
	"sync"
)

func sendData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func getData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Println("Received:", data)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go sendData(ch, &wg)
	go getData(ch, &wg)
	// Here instead of doing time.Sleep I am using waitGroup so that the program waits for all goroutines to complete before exiting.
	wg.Wait()
}
