package main

import("fmt")

func main() {

    p1 := struct {
      first string
      last string
      age int
    }{
      first: "James",
      last: "Name",
      age: 32 ,
    }

    fmt.Println(p1)
}
