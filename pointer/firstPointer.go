package main

import("fmt")

func main() {
  a := 42
  fmt.Println(a)
  fmt.Println(&a) // get the memory address
  fmt.Printf("%T\n", a)
  fmt.Printf("%T\n", &a)
  // var b *int = &a
  b := &a
  fmt.Println(*b) // dereference of the address ( which will get the value of a) we declare before
  fmt.Println(a) 
  fmt.Printf("%T\n", &b)
  *b = 44
  fmt.Println(a)  // this'll be 44 because a and b is sharing the same address
}
