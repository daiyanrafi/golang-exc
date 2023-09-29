package main

import "fmt"

func main () {
	var decimalNumber int

	fmt.Printf("Enter a decimal number: ")
	//scan dile %v deya lagto na
	fmt.Scanf("%v", &decimalNumber)

	fmt.Printf("your binary num %b\n", decimalNumber)
	fmt.Printf("your hex num %x\n", decimalNumber)
	fmt.Printf("your octal num %o\n", decimalNumber)
}