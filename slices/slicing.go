package main

import("fmt")

func main() { 
  //
  x := []int{10, 20, 30, 40, 50, 60} // x is a slice ([]) of int
  fmt.Println(x[1:2]); // output will be 20 because will get the 
                       // element that start from index 1 up to 
                       // element that before index 2 which is 20  
  fmt.Println(x[:2])   // start the element from 0 up until index 2
  fmt.Println(x[3:])   // start the element from 3 up until end index 

  for index, value := range x {
    fmt.Println (index, value)
  }

  for i := 0; i < 2 ; i++ {
    fmt.Println(i, x[i])
  }
}
