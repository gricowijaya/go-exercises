// - create a func with the identifier foo that   
//   * takes a variadic parameter of type int
//   * pass in a value of type []int into your func (unfurl the[]int)
//   * returns the sum of all values of type int passed in   

// - create a func with the identifier bar that   
//   * takes in a parameter of type []int
//   * returns the sum of all values of type int passed in   

package main 

import("fmt")

// create the slice of int from the variadic parameter 
func foo(args ...int) int { 
  var sum int;
  for _, value := range args { 
    sum += value;
  }
  return sum
}

// create the function that takes an argument of slice of int
func bar(x [] int) int { 
  var sum int;
  for _, value := range x { 
    sum += value;
  }
  return sum
}

// main function
func main() { 
  x := foo(1, 3, 4, 4);
  y := bar([]int{1,3,4,5});

  fmt.Printf("foo = %d\nbar = %d ", x, y);
}

