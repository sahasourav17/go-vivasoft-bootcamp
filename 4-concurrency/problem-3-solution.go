package main

import (
	"fmt"
	"sync"
)

func sendData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("Sending %d \n", i)
		ch <- i
	}
	close(ch)
}

func receiveData(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}

}
func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go sendData(ch, &wg)
	go receiveData(ch, &wg)

	wg.Wait()
}
