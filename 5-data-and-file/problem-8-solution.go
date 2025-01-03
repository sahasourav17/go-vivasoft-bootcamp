package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Create a map of data
	data := map[string]interface{}{
		"name":    "Raisul",
		"age":     23,
		"country": "BD",
	}
	// Create a new JSON file
	file, err = os.Create("dummy.json")
	if err != nil {
		fmt.Printf("Error creating JSON file: %v\n", err)
		return
	}
	defer file.Close()
	// Encode the data to JSON and write it to the file
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		fmt.Printf("Error encoding JSON data: %v\n", err)
		return
	}
	fmt.Println("JSON file generated successfully.")
}
