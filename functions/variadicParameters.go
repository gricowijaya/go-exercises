package main

import "fmt"

func main() {
  x := foo(1, 2, 3);
  fmt.Println("The total is ", x);
}

// foo is a variadic function. 
// It can be called in the following ways:
// foo()
func foo(args ...int) int { // use the lexical element
  fmt.Println(args)
  fmt.Printf("%T\n", args)

  sum := 0  // is more readable

  for i, r := range args {
    sum += r 
    fmt.Println("for item in index position", i, " We are adding ", r, " To the s total:", sum);
  }

  // fmt.Println("The total is ", sum)
  return sum
}
