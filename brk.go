package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue //ai condition eh asha manei abar loop eh fire jaoya nichey jabe na
			// break  //answer will be 1
		}
		fmt.Printf("%v\n",i)
	}
}	