// defer keyword to show that a deferred func runs 
// after the func containing it exists.
// defer function won't execute until another function has returns
package main

import("fmt")

func foo(x int) int {
  // after all foo is executed this'll be called
  defer func(){
    fmt.Println("this is the deffered function of foo and this is anonymous function")
  }()
  y := 100
  x *= y
  fmt.Println("Deffered foo value :", x)
  return  x
}

func main() {
  // 2. then after printing "Whatsapp" it'll call this
  defer foo(1)

  // 3. it'll execute after the first function executed and return a value
  defer foo(3)

  // 1. it'll call this first
  fmt.Println("Whatsapp")
}
  



