// create an anonymous function

package main

import("fmt")

func main() {
  x := []int32{1,2,3,4,5,6}

  func(){ 
    var sum int32;
    for _, values := range x {
      sum -= values
    }
    fmt.Print("This use the anonymous function\n");
    fmt.Printf("%d\n", sum);
  }()

  fmt.Println("this is outside the anonymous function")
}
