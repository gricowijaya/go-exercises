package main

import("fmt")

// if you have the method speak they also a type of human 
type human interface {
  speak()
} 

// something asking to an empty interface just like below
// because a value can be more then 1 type
// type human interface {
// } 

type person struct {
  first string
  last string
  phone string
}

type teacher struct {
  person
  isTeacher bool
}

func (p person) speak() {
  word := "Hallo last my name is" + p.last
  fmt.Println(word);
}

func (p teacher) speak() {
  word := "Hallo i'm a teacher and my last name is" + p.last
  fmt.Println(word);
}

// function to asserting the function
func bar (h human) {
  switch h.(type) {
    case person:
      fmt.Println("I was passed into bar my name is ", h.(person).first, " and i'm just a person")
    case teacher:
      fmt.Println("I was passed into bar my name is ", h.(teacher).first, " and i'm just a teacher")
  }
} 

func main() {
  var x person;
  var h teacher;
  x.first = "Tyas"
  h.first = "Tatang"
  bar(x)
  bar(h)
}

