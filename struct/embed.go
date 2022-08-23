package main
import("fmt")

type person struct {
  first string
  last string
}

type secretAgent struct {
  person
  ltk bool // license to kill
}

func main() {

    sa1 := secretAgent{
      person: person {
        first: "James",
        last: "Bond",
      },
      ltk: true,
    }
    
    p2 := person {
      first: "Miss",
      last: "Moneypenny",
    }

    fmt.Println(sa1.first, sa1.ltk)
    fmt.Println(p2)
}
