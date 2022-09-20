package main

import("fmt")

// using recursive function 
func factorialRecursive(n int) int {
  fmt.Println("Process ", n, " * ", n - 1)
  // stop the function for calling itself after we hit n == 0
  if n == 0 { return 1} 
  return n * factorialRecursive(n - 1) // return the factorial by calling it self
}

// using loops
func factorialLoop(n int) int {
  var x int
  // initialise the value x into 1 why ? because we will multiply the x with n
  // if we not assign x into 1 then we won't get the value.   
  x = 1 

  // initialize the incremental variable (i) then assign n into variable i 
  // why ? because we will decrement the value of n after it multiply by itself.
  for i := n; n > 0 ; n-- {
    x *= n 
    fmt.Printf("i[%d] x[%d] n[%d]\n", i, x, n) // trace the value on i, x, and n
  }

  // return the value of x which is the result of factorial of n
  return x
}

func main() {
  n := factorialRecursive(5);
  fmt.Println("The factorialRecursive value of 5 is = ", n)
  x := factorialLoop(5);
  fmt.Println("The factorialLoop value of 5 is = ", x)
}
