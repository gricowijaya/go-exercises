package main

import("fmt")

func main() { 
  //
  x := []int{1 , 2, 3, 4, 5, 6} // x is a slice ([]) of int
  fmt.Println(x);

  for index, value := range x {
    fmt.Printf("This is the element of [%d] is %d\n", index, value)
  }
}
