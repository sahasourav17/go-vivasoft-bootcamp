package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// Create a slice of Person structs
	people := []Person{
		{Name: "John Doe", Age: 30, Country: "USA"},
		{Name: "Jane Smith", Age: 25, Country: "Canada"},
	}
	// Create a new CSV file
	file, err := os.Create("people.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a writer to write to the file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	header := []string{"Name", "Age", "Country"}
	writer.Write(header)

	// Write the data rows from the slice to the CSV file
	for _, person := range people {
		row := []string{person.Name, strconv.Itoa(person.Age), person.Country}
		writer.Write(row)
	}

	fmt.Println("CSV file created successfully.")
}
