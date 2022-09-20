package main 

import (
  "fmt"
)

// built with iota
const ( 
  _ = iota;
  kb = iota * 1
  mb = iota * kb
  gb = iota * mb
) 

func main() {

  kb := 1024
  mb := 1024 * kb
  gb := 1024 * mb

  fmt.Println("\t[decimal]\t\t[binary]");
  fmt.Printf("[kb] == %d\t\t%b\n", kb, kb);  
  fmt.Printf("[mb] == %d\t\t%b\n", mb, mb);  
  fmt.Printf("[gb] == %d\t%b\n", gb, gb);    
}

