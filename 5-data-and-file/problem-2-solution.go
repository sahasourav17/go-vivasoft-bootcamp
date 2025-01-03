package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Country string   `json:"country"`
	Hobbies []string `json:"hobbies"`
	Gender  string   `json:"gender"`
}

func main() {
	str := `{
		"name": "John",
		"age": 20,
		"country": "USA",
		"hobbies": ["reading", "painting"],
		"gender": "male"
	}`

	res := response{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("Parsed Struct: %+v\n", res)

}
