package main

import ("fmt")

type Integer int

func main() { 
  var i int;
  i = 1;
  fmt.Printf("%T",i); // output :int
  fmt.Printf("%T",i); // output :string
}
