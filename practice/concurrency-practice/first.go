package main

import (
    "fmt"
    "sync"
    "runtime"
)

var wg sync.WaitGroup

func main() { 
    fmt.Println(runtime.NumGoroutine())
    fmt.Println(runtime.NumCPU())
    wg.Add(2);

    go firstRoutine();
    go secondRoutine()

    fmt.Println(runtime.NumGoroutine())
    fmt.Println(runtime.NumCPU())

    wg.Wait()

    fmt.Println(runtime.NumGoroutine())
    fmt.Println(runtime.NumCPU())
}

func firstRoutine() { 
    fmt.Println("First Routine: ")
    // for i := 1; i <= 30; i++  { 
    //     fmt.Println(i)
    // }
    wg.Done()
}

func secondRoutine() { 
    fmt.Println("second Routine: ")
    // for i := 0; i < 30; i++  { 
    //     fmt.Println(i)
    // }
    wg.Done()
}

