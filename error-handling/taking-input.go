package main

import(
    "fmt"
)

func main() { 
    var answer string 

    _, err := fmt.Scan(&answer);
    
    if err != nil { 
        panic(err)
    }

    fmt.Println(answer)
}
