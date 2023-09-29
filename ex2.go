package main

import "fmt"


func main () {

	// var fullName string = "John Doe"
	// var age int = 32
	// var gpa float32 = 3.14
	// var isCool bool = true

	// var fullName = "John Doe"
	// var age = 32
	// var gpa  = 3.14
	// var isCool = true

	fullName := "John Doe"
	age := 32
	const COUNTRY = "USA"
	gpa  := 3.14
	isCool := true


	fmt.Println(fullName, "is", age, "years old with a gpa of", gpa, "and is cool?", 
    isCool, "and lives in", COUNTRY)
}