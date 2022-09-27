package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type person struct {
    First string
    Last string
    Age int
}


func main() {
    // the data
    data := `[{"First":"Testing34","Last":"1","Age":32},{"First":"Testing","Last":"2","Age":33}]`

    // create the file path
    file := "./json/file.json"

    // create the file in here.
    f, err := os.Create(file)
    if err != nil {
        log.Fatal(err)
    }

    // create the buffer pass the file into the buffer for writing the data
    buffer := bufio.NewWriter(f)

    // get the value of the people slice value
    _, err = buffer.WriteString(data)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("this is the content of string literal json data :", file)

    // flush the buffer data to get cleaner buffer if we want to write again
    if err := buffer.Flush(); err != nil {
        log.Fatal(err)
    }
}
