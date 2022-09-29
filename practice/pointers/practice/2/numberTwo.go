/*
1. Create a person struct
2. Create a func called "changeMe" with a person as a parameter
   2.1 change the value stored at the *person address
3. create a value of type person
   3.1 print out the value
4. in func main   
   4.1 print out the value
5. important   
   5.1 to dereference a struct, use (*value).field
   5.2 as an exception, if the type of x is a named pointer type and (*x).f 
       is a valid selector expression denoting a field (but not a method), x.f 
       is shorthand for(*x).f

*/
package main

import("fmt")

// create the Person struct
type Person struct {
    name string
}

// function to change the value taking two parameter the pointer of the Person and string
func changeMe(p *Person, newName string) string   { 
    p.name = newName 
    return (*p).name
}

func main() {
    joko :=  Person {
        name: "jojo",   // assign the value of variable joko into jojo
    }

    fmt.Println("before", joko.name) // output : before jojo

    changeMe(&joko, "joko"); // here we pass the reference of joko variable

    fmt.Println("after", joko.name)  // output : after joko
}
