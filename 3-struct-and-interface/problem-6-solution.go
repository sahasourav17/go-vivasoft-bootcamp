package main

import "fmt"

type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) float64
}

type BasicCalculator struct{}

func (bc BasicCalculator) Add(a, b float64) float64 {
	return a + b
}

func (bc BasicCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (bc BasicCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (bc BasicCalculator) Divide(a, b float64) float64 {

	// if b == 0 {
	// 	fmt.Println("Error: Division by zero")
	// 	return 0
	// }
	// return a / b
	if b == 0 {
		panic("Cannot divide by zero")
	}
	return a / b
}

type ScientificCalculator struct{}

func (ac ScientificCalculator) Add(a, b float64) float64 {
	return a + b
}

func (ac ScientificCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (ac ScientificCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (ac ScientificCalculator) Divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func main() {
	// Create instances of calculators
	basicCalc := BasicCalculator{}
	scientificCalc := ScientificCalculator{}

	// Perform operations with BasicCalculator
	fmt.Println("Basic Calculator:")
	fmt.Printf("Add: %.2f\n", basicCalc.Add(10, 5))
	fmt.Printf("Subtract: %.2f\n", basicCalc.Subtract(10, 5))
	fmt.Printf("Multiply: %.2f\n", basicCalc.Multiply(10, 5))
	fmt.Printf("Divide: %.2f\n", basicCalc.Divide(10, 5))

	// Perform operations with ScientificCalculator
	fmt.Println("\nScientific Calculator:")
	fmt.Printf("Add: %.2f\n", scientificCalc.Add(12, 8))
	fmt.Printf("Subtract: %.2f\n", scientificCalc.Subtract(12, 8))
	fmt.Printf("Multiply: %.2f\n", scientificCalc.Multiply(12, 8))
	fmt.Printf("Divide: %.2f\n", scientificCalc.Divide(12, 8))
}
