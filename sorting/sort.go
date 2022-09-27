package main

import(
    "fmt"
    "sort"
)

func main() {
    sliceOfInt := []int{4, 7, 3, 42, 18, 16, 56, 12}
    sliceOfString := []string{"B", "S", "A", "E", "U", "N", "C", "O", "R"}

    fmt.Println("before", sliceOfInt)          // output : [4 7 3 42 18 16 56 12] 
    fmt.Println("before", sliceOfString)       // output : [4 7 3 42 18 16 56 12] 

    // call the function according the data type of the slice
    // this function below will sort the slice in increasing order
    sort.Ints(sliceOfInt)
    sort.Strings(sliceOfString)

    fmt.Println("after", sliceOfInt)           // output : [4 7 3 42 18 16 56 12] 
    fmt.Println("after", sliceOfString)        // output : [4 7 3 42 18 16 56 12] 
}
