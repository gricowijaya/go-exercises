package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

    // set the counter or incrementer
	var counter int64 // which is a really big flag that they use.


	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
            // we want to add the counter and just add 1 
			atomic.AddInt64(&counter, 1)
            fmt.Println("Counter in this is ", atomic.LoadInt64(&counter)); // to print it out we need the lock

			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
