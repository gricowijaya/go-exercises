package main

import(
    "fmt"
    "sort"
)

type person struct {
    first string
    age int
}

// to do custom sort we can 
// 1. create a type call byAge which is a person type. 
type ByAge []person

// 2. then create a function that'll accepting ByAge type 
//    type for the implementation of ByAge sorting
// each function to get the length, swapping the value, 
// and checking either the value is lesser or greater 
// then the next value  
// and the function below is implementing the Sort method
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
    // define the person
    p1 := person{"Ajeng", 20}
    p2 := person{"Dimas", 32}
    p3 := person{"Zul", 40}
    p4 := person{"Wahyu", 25}

    // create a slice of person which resulting in people 
    people := []person{p1, p2, p3, p4}

    fmt.Println("before sort by age ", people)

    sort.Sort(ByAge(people));

    fmt.Println("after sort by age", people)
}
