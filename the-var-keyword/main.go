package main

import ("fmt")

var y int;

func main() {
  y = 102;
  fmt.Printf("The Hexadecimal of %d is %#x meanwhile in binary is %b\n", y, y, y); // take the argument of every format specifier
  fmt.Println("Here can try to change the y variable from 102 into 66"); // take the argument of every format specifier
  y = 19; // reassign
  fmt.Printf("The Hexadecimal of %d is %#x meanwhile in binary is %b\n", y, y, y); // take the argument of every format specifier
}
