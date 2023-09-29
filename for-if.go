// while wont work in go but we can modify it to for loop
package main

import "fmt"

func main() {

	// for x:=2; x<=100 ; x+=2 {
	// 	fmt.Printf("%v\n",x)
	// }

	for x:=2; x<=100 ; x+=2 {
		if x % 2 == 0 {
			fmt.Printf("%v\n",x)
		}
	}
}

