package main

import (
	"fmt"

	"golang.org/x/text/message"
)

func displayGG(home int) {
	result := home * home
	fmt.Printf("i love %v\n", result)
}

func add (num1 int, num2 int) int{
	result := num1 + num2
	return result;
}

func sub (num1 int, num2 int){
	result := num1 - num2
	fmt.Printf("%v\n", result)
}

func msg() string {
	return "GG hello"
}

func main() {
	displayGG(5)
	// add(10,20)
	sub(10,10)
	fmt.Printf("%v\n",add(10,20))
	fmt.Printf(msg())
}