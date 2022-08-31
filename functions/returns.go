package main
import("fmt")

func main() {
  stringThis := foo()
  fmt.Println(stringThis)

  x := bar()  
  fmt.Printf("%T\n", x)       // going to return a func int() which is the type of x    
  fmt.Printf("%d\n", x())     // return 451
  
}

func foo() string {
  s := "Hello World"
  return s
}

// function bar is a will return a function that returns an intkj
func bar() func() int {
  return func() int {
    return 451
  }
}
