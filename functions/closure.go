package main

import("fmt")

func main() {
  i := incrementor() // i variable is a function expression of incrementor()
  j := incrementor() // j variable is a function expression of incrementor()

  fmt.Println("i func expression", i()) // output : 1
  fmt.Println("j func expression", j()) // output : 1  
  fmt.Println("j func expression", j()) // output : 2 (x + 1)
  fmt.Println("j func expression", j()) // output : 3 (x + 1)
  fmt.Println("j func expression", j()) // output : 4 (x + 1)
  fmt.Println("j func expression", j()) // output : 5 (x + 1)
  fmt.Println("j func expression", j()) // output : 6 (x + 1)
}


// this function will return a function that return an int
func incrementor() func() int {
  // this `x` variable below will be valid until a few line
  // each time this function is called then golang will 
  // create a new address for the x variable         
  // that's why even if this is a local function it still 
  // can increment the value of the value of x 
  var x int 
  return func() int{
     x++ // increment the x + 1;
     return x
  }
}

