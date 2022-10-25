package main

import (
	"fmt"
	"log"
	"os"
)

// read data
func main() {
    _, err := os.Open("file.txt")

    if err != nil { 
        fmt.Println("there's been some error -> ", err) 
        log.Println("err, happened", err)
        log.Fatal(err)
        panic(err)
    }
}
