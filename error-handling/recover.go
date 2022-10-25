package main

import ("fmt")

func deferred() { 
    fmt.Println("\nDeferred")
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}

func not_deferred(){ 
    fmt.Println("\nNot Deferred")
    for i := 0; i < 4; i++ {
        fmt.Print(i)
    }
}

func main() { 
    deferred();
    not_deferred();
}

