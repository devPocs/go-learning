package main

import "fmt"

//function that returns two values
	func multiply (a, b int) (int, string) {
		result := a * b
		return result, "Successful multiplication"
	}
	//fuction that returns three values
	func person (name string) (string, int, bool) {
		fullName:= "Mr. "+ name
		age:=30
		isEmployed:= true

		return fullName, age, isEmployed
	}

func main() {
result, message := multiply(2, 15)
fmt.Printf("The result of our multiplication is: %d. %s.", result, message )

fullName, age, isEmployed := person("John Bellow")
fmt.Printf(("\nFull Name: %s, Age: %3d, Is Employed? %t"), fullName, age, isEmployed)
}