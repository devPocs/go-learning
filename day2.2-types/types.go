package main

import "fmt"
func main(){
//Go basic types

//string types
var name string = "Johnny Depp"

//integer types 
var smallInt int8 = 100
var regularInt int = 1000
var bigInt int64= 1234567890

//unsigned integer types- only positive values
var positiveOnly uint = 25 

//float types
var smallFloat float32 = 3.142
var preciseFloat float64 = 3.141592653589793

//boolean type
var isLearning bool = true

//character type
var character byte = 'A'

fmt.Printf("Name: %s. (I am of Type :%T)", name, name)
fmt.Printf("\nSmall Integer: %d. (I am of Type: %T)", smallInt, smallInt)
fmt.Printf("\nRegular Integer: %d. (I am of Type: %T)", regularInt, regularInt)
fmt.Printf("\nBig Integer: %d. (I am of Type: %T)", bigInt, bigInt)
fmt.Printf("\nPositve Only Integer: %d. (I am of Type: %T)", positiveOnly, positiveOnly)
fmt.Printf("\nSmall Float: %3f. (I am of Type: %T)", smallFloat, smallFloat)
fmt.Printf("\nPrecise Float: %f. (I am of Type: %T)", preciseFloat, preciseFloat)
fmt.Printf("\nBoolean: %t. (I am of Type: %T)", isLearning, isLearning)
fmt.Printf("\nCharacter: %c. (I am of Type: %T)", character, character)
}
