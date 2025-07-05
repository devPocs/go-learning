package main

import "fmt"

func main() {
	var name string
	var num1, num2 float64

	// get user input
	fmt.Println("Please, enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Printf("hello %s! Here is a simple calculator:\n\n", name)
	

	fmt.Printf("%.1f + %.1f = %.1f\n", num1, num2, num1+num2)
	 fmt.Printf("%.1f - %.1f = %.1f\n", num1, num2, num1-num2)
    fmt.Printf("%.1f × %.1f = %.1f\n", num1, num2, num1*num2)

	if num2 !=0 {fmt.Printf("%.1f ÷ %.1f = %.4f\n", num1, num2, num1/num2)
} else{fmt.Println("Division by zero is not allowed.", num2 , "can't be used as a divisor.")}

    
}