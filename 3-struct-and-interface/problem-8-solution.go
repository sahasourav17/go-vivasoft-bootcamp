package main

import (
	"fmt"
	"strings"
)

// Formatter interface
type Formatter interface {
	Format(data []string) string
}

// CSVFormatter struct
type CSVFormatter struct{}

// Format method for CSVFormatter
func (c CSVFormatter) Format(data []string) string {
	return strings.Join(data, ",")
}

// HTMLFormatter struct
type HTMLFormatter struct{}

// Format method for HTMLFormatter
func (h HTMLFormatter) Format(data []string) string {
	html := "<ul>\n"
	for _, item := range data {
		html += "  <li>" + item + "</li>\n"
	}
	html += "</ul>"
	return html
}

// TestFormat function
func TestFormat(f Formatter, data []string) {
	// Call Format method
	fmt.Println("Formatted Output:")
	fmt.Println(f.Format(data))

	// Use reflection to determine the type of the formatter
	switch f.(type) {
	case CSVFormatter:
		fmt.Println("The formatter is a CSV formatter.")
	case HTMLFormatter:
		fmt.Println("The formatter is an HTML formatter.")
	default:
		fmt.Println("Unknown formatter type.")
	}
}

func main() {
	// Test CSV formatter
	csvFormatter := CSVFormatter{}
	TestFormat(csvFormatter, []string{"apple", "banana", "orange"})
	fmt.Println()

	// Test HTML formatter
	htmlFormatter := HTMLFormatter{}
	TestFormat(htmlFormatter, []string{"apple", "banana", "orange"})
}
