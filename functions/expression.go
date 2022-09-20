package main

import("fmt")

func main() {
  foo()

  // this is another anonymous function with no argument passed 
  func() {
    fmt.Println("Anonymous func")
  }() // include the parenthses which pass nothing

  // this is another anonymous function with argument passed 
  print42 := func(x int) {
    fmt.Printf("Anonymous func which take args of x = %d", x);
  }(42) // include the parenthses which pass nothing

  // because we need to output of the value of 42 then this print 42   
  print42(); }

func foo() {
  fmt.Println("foo() Function Hello World")
}
