package main

import (
  "fmt"
)

type person struct {
  first string
  last string

}

type secretAgent struct {
  person
  ltk bool
}

// any value with the the type of secret Agent have the method of speak
func (s secretAgent) speak() {
  fmt.Println("I am ", s.first, s.last, " and I have the licencse to kill value of ", s.ltk);

}

func main() {
  // value of secretAgent type have the access the method of speak
  firstSecretAgent := secretAgent {
    person: person {
      "James", 
      "Bond",
    },
    ltk:true,
  }
  fmt.Println(firstSecretAgent)
  firstSecretAgent.speak();
}
