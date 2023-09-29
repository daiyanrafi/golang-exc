package main

import "fmt"

func main() {

	var endNum int
	var srtNum int

	fmt.Printf("enter start num : ")
	fmt.Scan(&srtNum)

	fmt.Printf("enter last num : ")
	fmt.Scan(&endNum)

	sum:=0

	for i := srtNum; i <= endNum; i++ {
		sum = sum + i
	}
	fmt.Printf("sum = %v\n",sum)
}	