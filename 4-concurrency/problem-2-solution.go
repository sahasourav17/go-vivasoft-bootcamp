package main

import (
	"fmt"
	"sync"
)

func calculateSquare(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Square of %d is %d\n", num, num*num)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go calculateSquare(ch, &wg) // Starting the goroutine

	for {
		var input int
		fmt.Print("Enter an integer (or -1 to exit): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid integer.")
			continue
		}

		if input == -1 {
			break
		}

		ch <- input
	}

	close(ch) // closing the channel
	wg.Wait() // Wait for the  goroutine to finish
	fmt.Println("Program terminated.")
}
