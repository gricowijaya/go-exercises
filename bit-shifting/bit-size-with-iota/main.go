package main 

import (
  "fmt"
)

// built with iota
const ( 
  _ = iota;
  kb = 1 << ( iota * 10  )
  mb = 1 << ( iota * 10 )
  gb = 1 << ( iota * 10 )
) 

func main() {
  fmt.Println("\t[decimal]\t\t[binary]");
  fmt.Printf("[kb] == %d\t\t%b\n", kb, kb);  
  fmt.Printf("[mb] == %d\t\t%b\n", mb, mb);  
  fmt.Printf("[gb] == %d\t%b\n", gb, gb);    
}

