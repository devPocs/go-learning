package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	} else {
		return a / b, nil
	}
}

func ageValidation(age int) (string, error) {
	if age <= 0 {
		return "", errors.New("age cannot be negative or equal to zero")
	}
	if age > 150 {
		return "", errors.New("age is unrealistic")
	}
	return "Valid age", nil
}

func main() {
	fmt.Println("Error Handling in Go")
	result, err := divide(10, 0)

	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("Result: %.2f", result)
	}

	message, err := ageValidation(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Message:", message)
	}
}
