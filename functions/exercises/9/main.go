package main

import("fmt")

func callbackExercices(f func(xi []int, ii []int) int)  int {
  n := f(ii);
  n++
  return n;
  // x:= callbacked(); 
  // fmt.Print("The callbacked function just like this", x);
}

func main() {
  g:= func(xi []int) int{
    if len(xi) == 0 {
      return 0;
    }
    if len(xi) == 1 {
      return xi[0];
    }
    return xi[0] + xi[len(xi) - 1];
  }

  callbackExercices(g);
}
