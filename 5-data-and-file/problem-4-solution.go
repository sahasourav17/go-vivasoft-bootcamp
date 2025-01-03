package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	data := []byte("Hello, world!")
	n, err := file.Write(data)
	if err != nil {
		fmt.Printf("File write error: %v\n", err)
		return
	}

	fmt.Printf("Wrote %d bytes to file\n", n)
}
