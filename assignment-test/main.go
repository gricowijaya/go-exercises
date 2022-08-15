package main 

import("fmt")

func main() {
  a := ( 42 == 42) // output : true
  b := ( 42 <= 43) // output : true
  c := ( 42 >= 43) // output : false
  d := ( 42 != 43) // output : true
  e := ( 42 <  43) // output : true
  f := ( 42 >  43) // output : false

  fmt.Println(a, b, c, d, e, f);
}

