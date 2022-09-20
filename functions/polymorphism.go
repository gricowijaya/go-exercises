package main

import (
	"fmt"
	"os"
)

type person struct {
  first string
  last string
}

type secretAgent struct {
  person
  ltk bool
}


// anybody or type who have the method of speak is implement an human interface   
// keyword identifier type
type human interface {
  speak()
}

func bar(h human) {
  fmt.Println("i called human ", h);
}

//---------------------------------------------
// BELOW is the same method but different type

// any value with the the type of secret Agent have the method of speak
func (s secretAgent) speak() {
  fmt.Println("I am ", s.first, s.last, " and I have the licencse to kill value of ", s.ltk);

}

// let's add the person to the method for speak so it's implementing the human interface
func (p person) speak() {
  fmt.Println("I am ", p.first, p.last, " and I am a person");
}

// create a new type called hotdog
type hotdog int

func main() {
  firstSecretAgent := secretAgent {
    person: person {
      "James", 
      "Bond",
    },
    ltk:true,
  }

  firstPerson := person {
    first: "Dr. ", 
    last: "Octopus",
  }

  fmt.Println(firstSecretAgent);
  fmt.Println(firstPerson);
  firstSecretAgent.speak();

  // first Person is a type of person but it has a function of person 
  // so we can pass the parameter of first person into a human parameter.  
  bar(firstPerson); 
  
  // conversion use the type of hotdog
  var x hotdog = 42
  fmt.Println(x)
  y := int(x) // conversion around hotdog

  fmt.Println("type of x %T\n", x);
  fmt.Println("type of y %T\n", y);
}

