package main 

import (
  "fmt"
)

func main() {
  x := 2
  fmt.Printf("[x] == %d\t\t%b\n", x, x); // output : %d is short for decimal and %b is for binary 

  y := x << 1 //  assign shited x by 1 into y
  fmt.Printf("[x] == %d\t\t%b\n", y, y); // output : %d is short for decimal and %b is for binary 
}

