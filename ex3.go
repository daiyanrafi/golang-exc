package main

import "fmt"

func main () {
	var fullName string
	var age int
	var gpa float32

//input: scan, scanln,scanf
//output: print, println, printf, printfln

fmt.Print("Enter your full name: ")
fmt.Scan(&fullName)

fmt.Print("Enter your full age: ")
fmt.Scan(&age)

fmt.Print("Enter your full gpa: ")
fmt.Scan(&gpa)

fmt.Printf("Hello %v\n", fullName)
fmt.Printf("You are %v years old and gpa %v", age, gpa)

}