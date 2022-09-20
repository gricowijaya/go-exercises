package main

import("fmt")

func main() { 
  s := "Hello World"
  fmt.Println(s);
  fmt.Printf("%T\n", s);

  bytesString := []byte(s); // print the slice of bytes the letters decimal 
  fmt.Println(bytesString);
  fmt.Printf("%#x", bytesString); // print in the hexadecimal
  fmt.Printf("%T\n", bytesString); // alias of uint8
}
