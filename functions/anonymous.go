package main

import("fmt")

func main() {
  foo()

  // this is another anonymous function with no argument passed 
  func() {
    fmt.Println("Anonymous func")
  }() // include the parenthses which pass nothing

  // this is another anonymous function with argument passed 
  func(x int) {
    fmt.Printf("Anonymous func which take args of x = %d", x);
  }(42) // include the parenthses which pass nothing
}

func foo() {
  fmt.Println("foo() Function Hello World")
}
