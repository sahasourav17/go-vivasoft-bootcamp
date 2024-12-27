package main

import (
	"fmt"
)

func PrintValues(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}
func main() {
	vals := []interface{}{1, "hello", true}
	PrintValues(vals)
}
