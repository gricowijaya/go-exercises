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
    data := `[{"First":"Testing","Last":"1","Age":32},{"First":"Testing","Last":"2","Age":33}]`
    byteOfString := []byte(data) 

    var people []person

    fmt.Println("this is the content of string literal json data :", data)

    // unmarshal the data and save into the people memory address
    err := json.Unmarshal(byteOfString, &people)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("this is the content of unmarshaled data in the people struct:", people)

    // get the value of the people slice value
    for i, value := range people {
        fmt.Println("The Index ", i, " value of ", value)
    }  
}
