package main

import ("fmt")

func main() {
  foo();
  bar("James")
  fmt.Println(woo("Rico"))
  x, y := mouse("Ian", "Fleming")
  fmt.Println(x)
  fmt.Println(y)
}

func foo() {
  fmt.Println("Hello from Foo!")
}

// pass by value
func bar(s string) {
  fmt.Println("Hello,", s);
}

// return value of string
func woo(st string) string { // becuase we need to return a stirng {
  return fmt.Sprint("Hello form Woo, Your name is ", st)
}
// return string and bool 
func mouse(fn string, ln string) (string, bool) {
  a := fmt.Sprint(fn, ln, `, says "Hello"`)
  b := true

  return a, b
} 
