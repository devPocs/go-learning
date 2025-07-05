package main

import "fmt"

func main() {
	//method 1: using var keyword and explicit type declaration
	var name string = "Pokoh Ufuoma"
	var age int = 25

	//method 2: var with type inference
	var city = "New York"

	//method 3: short declaration
	actor_name := "Dwayne Johnson"

	//method 4: multiple varibles declaration at once

	var movie1, movie2 string = "Mummy's Return", "Fast and Furious"

	fmt.Printf("My name is %s, I am %d years old. I live in %s. My favorite actor is %s.\nMy favorite movies are The %s and The %s.\nThanks!", name, age, city, actor_name, movie1, movie2 )
}