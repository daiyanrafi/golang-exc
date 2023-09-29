package main

import "fmt"

func main() {
    var base, height, area float32

    fmt.Printf("Enter the base of the triangle: ")
    fmt.Scan(&base) // Use %f to read a float32 value

    fmt.Printf("Enter the height of the triangle: ")
    fmt.Scan( &height) // Use %f to read a float32 value

    area = 0.5 * base * height

    fmt.Printf("Area is = %v%%\n", area)
}
