package main

import(
    "fmt"
    "runtime"
    "sync"
)

func main() {
    // the go max procs is set according to the core you have 
    fmt.Println("$CPU", runtime.NumCPU())
    fmt.Println("$GoRoutines", runtime.NumGoroutine())

    counter := 0
    const gs = 100

    var wg sync.WaitGroup
    var mu sync.Mutex

    wg.Add(gs)

    // use the GoSched
    for i:= 0; i < gs; i++ { 
        go func() { 
            // lock the variables
            mu.Lock()
            v := counter 
            // time.Sleep(time.Second) // put the sleep the 
            runtime.Gosched()
            v++
            counter = v // take the value and write it down into the shared variable and using  the runtime.Gosched()kj:w
            mu.Unlock()
            wg.Done()
        }()
        fmt.Println("$GoRoutines", runtime.NumGoroutine()) // we can see there's so many go routines goin ong
    }

    wg.Wait()
    fmt.Println("$CPU", runtime.NumCPU())
    fmt.Println("$GoRoutines", runtime.NumGoroutine())
    fmt.Println("$count", counter)
}
