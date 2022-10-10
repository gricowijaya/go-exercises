package main

import("fmt")

func main() { 
    eve:= make(chan int)
    odd:= make(chan int)
    quit:= make(chan int)


    // send
    go send(eve, odd, quit)

    // receive, we want to pull the value of the channels
    receive(eve, odd, quit)
    fmt.Println("About to exit")

}

func receive(e, o, q <- chan int) { 
    // loop until the condition of the select statement is satisfied
    for { 
        // select statement just like the switch case
        select { 
            // if the value is even then print
            case v := <- e:
                fmt.Println("From the eve channel ", v);
            case v := <- o:
                fmt.Println("From the odd channel ", v);
            case v := <- q:
                fmt.Println("From the quit channel ", v);
            return 
        }
    }

}

func send(e, o, q chan <- int) { 
    for i := 0; i < 100; i++ {
        if i % 2 == 0 {
            e <- i
        } else {
            o <- i
        }
        close(e)
        close(o)
        q <- 0
    }  
}
