package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if (i%3 == 0 || i%5 == 0) && i%15 != 0 {
            fmt.Println(i)
        }
    }
}


// Find all numbers between 1 and 100 that are divisible by 3 or 5, but not by 15 ?


