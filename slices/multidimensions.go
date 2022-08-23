package main

import("fmt")

func main() { 
  x := []int{10, 20, 30, 40, 50, 60}
  y := []int{1, 45, 433, 432, 678}

  z := [][]int{x, y}
  fmt.Println("multi dimensional slices ", z)
}
