package main

import("fmt")

func main() { 
  x := []int{10, 20, 30, 40, 50, 60} // x is a slice ([]) of int
  fmt.Println("before append", x)
  x = append(x, 80, 70, 90 )         // append 80, 70, 90 into x
  fmt.Println("after append", x)

  y := []int{1,45,433,432,678}
  x  = append(x, y...)

  fmt.Println("After appended with y ", x)

  x = append(x[:0], x[6:]...) // remove the element between the 0th and 6th index 
  fmt.Println("After remove the element ", x)
}
