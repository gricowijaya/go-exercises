package main

import (
    "fmt"
    "encoding/json" // import the json module
)

// we need to create the key to be uppercase
// before
// type user struct {
//     first string
//     age int
// } 
// after

// create the struct
type user struct {
    First string
    Age int
} 

func main() {
    u1 := user {
        First: "Dani",
        Age: 18,
    }

    u2 := user {
        First: "Rama",
        Age: 38,
    }

    u3 := user {
        First: "Ajeng",
        Age: 20,
    }

    // put the slice of user into users
    users := []user{u1, u2, u3} 
    
    // here we convert the slice of user into JSON
    bs, err := json.Marshal(users)
    if err != nil {
        fmt.Println(err) // error handling
    }

    // type cast the byte of string into string 
    fmt.Println(string(bs));
}
