package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully!")
}
