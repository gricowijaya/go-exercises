package main

import ("fmt")

func main() {
  ii := []int{1,3,4,1,5,6}
  fmt.Println("Hello, World");
  s := sum(ii...)
  s2 := even(sum, ii...) // pass the function of sum and also the ii slice
  s3 := odd(sum, ii...)  // check the function
  fmt.Println(s)
  fmt.Println("All even Numbers", s2)
  fmt.Println("All odd Numbers", s3)
}

func sum(xi ...int) int {
  fmt.Printf("%T\n", xi) // output : slice of int
  total := 0
  for _, value := range xi {
    total += value
  }
  return total
}

// passing a function of variadic parameter of int 
// sum of all even number
func even(f func(xi ...int) int, vi ...int) int{
  var yi []int
  for _, value := range vi {
    if value % 2 == 0 {
      yi = append(yi, value);
    }
  }

  return f(yi...) // getting the unilimited number of slice of int
}

// passing a function of variadic parameter of int 
// sum of all odd number
func odd(f func(xi ...int) int, vi ...int) int{
  var yi []int
  for _, value := range vi {
    if value % 2 != 0 {
      yi = append(yi, value);
    }
  }

  return f(yi...) // getting the unilimited number of slice of int
}
