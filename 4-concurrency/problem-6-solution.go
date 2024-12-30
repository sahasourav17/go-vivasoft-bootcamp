package main

import (
	"fmt"
)

func printEvenNumbers(ch chan int, done chan bool) {
	for i := 0; i <= 100; i += 2 {
		ch <- i
	}
	done <- true
}

func printOddNumbers(ch chan int, done chan bool) {
	for i := 1; i <= 100; i += 2 {
		ch <- i
	}
	done <- true
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	done := make(chan bool, 2) // To track completion of both goroutines

	// Start both goroutines
	go printEvenNumbers(evenCh, done)
	go printOddNumbers(oddCh, done)

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d  Even\n", <-evenCh)
		} else {
			fmt.Printf("%d  Odd\n", <-oddCh)
		}
	}

	// Wait for both goroutines to finish
	<-done
	<-done
}
