// create a user defined struct with
//   * the identifier person
//   * fields: first, last, age
// attach a method to type person with 
//   * the identifier "speak"
//   * the method should have the person say their name and age
// create a value of type person
//   call the method from the value of type person
  
package main

import("fmt")

type person struct{
  first string
  last string
  age int
}

func (p person) speak() {
  fmt.Print("My first name is ", p.first, " i am ", p.age, " and my last name is ", p.last)
}

func main() { 
  var Sarah person
  Sarah = person {
    "Sarah",
    "John",
    19,
  }

  Sarah.speak()
}
