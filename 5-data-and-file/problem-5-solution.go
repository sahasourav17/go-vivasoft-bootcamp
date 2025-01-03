package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("File content: %v\n", string(file))
}
