package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Country string   `json:"country"`
	Hobbies []string `json:"hobbies"`
	Gender  string   `json:"gender"`
}

func main() {
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Country: "USA",
		Hobbies: []string{"reading", "painting"},
		Gender:  "Male",
	}

	jsonString, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonString))
}
