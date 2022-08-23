package main
import("fmt")

type person struct {
  first string
  last string
  age int
}

func main() {
    p1 := person {
      first: "James",
      last: "Bond",
    }
    
    p2 := person {
      first: "Miss",
      last: "Moneypenny",
    }

    fmt.Println("First p1", p1.first)
    fmt.Println("Last  p2", p2.last)
}
