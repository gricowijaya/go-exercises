package main

import (
	"fmt"
	"runtime"
	"sync"
)

// run before the main function program usually using at the web dev.
// func init() {

// }

// wait for the go routines to finish
var wg sync.WaitGroup

// the sequential we will lanunc the go routines.
func main() {
	// getting the os of your create
	fmt.Println("OS", runtime.GOOS)
	// getting the create
	fmt.Println("ARCH", runtime.GOARCH)
	// getting the core of your computer cpu
	fmt.Println("Number of CPU", runtime.NumCPU())
	// getting the goroutines
	// our control flow won't need to have to wait the go routines.
	fmt.Println("Number of Go Routines", runtime.NumGoroutine()) // output : 1

	// call the function using the go routines keyword
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("Number of Go Routines", runtime.NumGoroutine()) // output : 2
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // remove that one thing
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
