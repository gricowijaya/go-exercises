package main

import("fmt")

func main() {
  for i := 0; i <= 3 ; i++ {
    if i == 3 {
      break;
    }

    fmt.Println("it's already broken in ", i);
  }
}
