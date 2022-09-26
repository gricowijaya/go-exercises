package main

import(
    "fmt"
    "encoding/json"
)

type person struct {
    First string
    Last string
    Age int
}


func main() {
    p1 := person {
        First :"Testing",
        Last: "1", 
        Age: 32,
    }

    p2 := person {
        First :"Testing",
        Last: "2", 
        Age: 33,
    }

    people := []person{p1, p2}
    
    bs, err := json.Marshal(people) 
    if err != nil { 
        fmt.Println(err)
    }
    fmt.Println(string(bs))
}
