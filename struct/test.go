package main

import("fmt")

var x, y int

type person struct {
  first string
  last string
  age int
}

type secretAgent struct {
  person // promoted fields which is unqualified
  ltk bool
}

func main() {
  x = 42
  y = 43
  fmt.Println(x, y)

  x, y := 44, 46
  fmt.Println(x, y);
}
