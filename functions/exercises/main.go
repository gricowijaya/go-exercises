package main

import("fmt")

func foo(x int) int { 
  var value int
  value = 1 + x
  return value
}

func bar() (int, string) { 
  return 1983, "Whatsapp"
}

func main() {
  x := foo(2)  // return 
  y, s := bar() // declare the y and s for storing the multiple return value. 

  fmt.Printf("%s %d\n", s, x);
  fmt.Printf("%d\n", y);
}

