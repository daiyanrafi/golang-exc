package main

import "fmt"

func change (val int) {
  val = 8
}

func GG (val *int) {
  *val = 8000
}

func main(){
  x:=10
  change(x)
  fmt.Println(x)

  // call by reference
  GG(&x)
  fmt.Println(x)
}