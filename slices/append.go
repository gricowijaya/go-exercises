package main

import("fmt")

func main() { 
  x := []int{10, 20, 30, 40, 50, 60} // x is a slice ([]) of int
  fmt.Println("before append", x)
  x = append(x, 80, 70, 90 )
  fmt.Println("after append", x)
  // _ := []string{ "ten", "twenty", "thirty", "fourty", "fifty", "sixty" }
  z := []int{ 132, 3, 1, 23, 35 }
  x = append(x, z...)
}
