package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Printf("Enter three numbers separated by space: ")
	fmt.Scan(&num1, &num2, &num3)

	if num1 > num2 && num1 > num3{
			fmt.Printf("num 3 is big %v\n", num1)
		} else if num2 > num1 && num2 > num3{
			fmt.Printf("num 3 is big %v\n", num2)
		} else if num3 > num1 && num3 > num2{
			fmt.Printf("num 3 is big %v\n", num3)
		}  else {
		fmt.Printf("num equal\n")
	}
}