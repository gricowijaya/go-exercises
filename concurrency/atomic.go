package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// the go max procs is set according to the core you have
	fmt.Println("$CPU", runtime.NumCPU())
	fmt.Println("$GoRoutines", runtime.NumGoroutine())

    // create the variable to run the int 64 and countering the int 64
    var counter int64

	// counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	// use the GoSched
	for i := 0; i < gs; i++ {
		go func() {
            // we can write and read in the safe environment
            atomic.AddInt64(&counter, 1)
            runtime.Gosched()
            fmt.Println("Counter", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("$GoRoutines", runtime.NumGoroutine()) // we can see there's so many go routines goin ong
	}

	wg.Wait()
	fmt.Println("$CPU", runtime.NumCPU())
	fmt.Println("$GoRoutines", runtime.NumGoroutine())
	fmt.Println("$count", counter)
}
