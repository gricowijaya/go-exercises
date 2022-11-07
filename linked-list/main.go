package main

import( 
    "fmt"
    "container/list"
)

func main() {
    // create a new list let's call it node
    Node := list.New()
    Node.PushBack(1)
    Node.PushFront(5)

    for i := 0; i < 10; i++ { 
        // if it's even put the element into back of the node
        if (i % 2 == 0) {
            Node.PushBack(i)
        } else {
            Node.PushFront(i)
        }
    }

    // print the element
    for element := Node.Front(); element != nil; element = element.Next() {
        if element.Value == 0 { 
            Node.Remove(element);
        }
    }

    for element := Node.Front(); element != nil; element = element.Next() {
        fmt.Println(element.Value)
    }
}

