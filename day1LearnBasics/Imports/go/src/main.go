package main

// import the modules

import (
  "fmt"
  "os"
)

func main() {
  // the arguments takes from array 0 and that's just to read the path of that directory.
  if len(os.Args) != 2 {
    println("Hello World with condition")
    os.Exit(1)
  }

  fmt.Println("Testing Arguments " + os.Args[1])
}

// documentation at godoc -http=:6060
