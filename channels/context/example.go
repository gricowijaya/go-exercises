package main

import(
    "fmt"
    "context"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    
    // defer cancel is cancel when we are finished

    for v := range gen(ctx) {
        fmt.Println(v)
        if

    }
}
