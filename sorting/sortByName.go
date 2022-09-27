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
type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
    // define the person
    p1 := person{"Ajeng", 20}
    p2 := person{"Dimas", 32}
    p3 := person{"Zul", 40}
    p4 := person{"Wahyu", 25}

    // create a slice of person which resulting in people 
    people := []person{p1, p2, p3, p4}

    fmt.Println("before sort by name ", people)

    sort.Sort(ByName(people));

    fmt.Println("after sort by name", people)
}
