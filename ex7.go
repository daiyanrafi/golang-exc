package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Printf("Enter three numbers separated by space: ")
	fmt.Scan(&num1, &num2, &num3)

	if num1 > num2 {
		if num1 > num3 {
			fmt.Printf(" num %v is larger\n", num1)
		} else {
			fmt.Printf("num 3 is big %v\n", num3)
		}
	} else if num2 > num1{
		if num2 > num3 {
			fmt.Printf("num 2 is big %v\n", num2)
		} else {
			fmt.Printf("num 3 is big %v\n", num3)
		}
	} else {
		fmt.Printf("num equal\n")
	}
}