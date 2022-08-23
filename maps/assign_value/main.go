package main

import("fmt")

func main() {
  array := [5]int{1,2,2,4,5,}

  for _, value := range array {
    fmt.Println(value)
  }
}
