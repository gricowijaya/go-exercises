package main

import("fmt")

func main() {
    x := 0
    fmt.Println("in main before value", &x); // output 0
    fmt.Println("in main before address",x);    // address of 0
    foo(&x)
    fmt.Println("in main before value", &x); // output 0
    fmt.Println("in main before address",x);    // address of 0

}

// we can pass the address we can mutate the value 
func foo(y *int) {
    fmt.Println("in foo before value", y); 
    fmt.Println("in foo before address",*y); // get the value of the address
    // dereference
    *y = 43
    fmt.Println("in foo before value", y);
    fmt.Println("in foo before address",*y);
}
