package main

import(
    "fmt" 
    "math"
)

type circle struct {
    radius float64
}

// method area is attached to shape type
type shape interface {
    area() float64
    area_with_no_pointer_receiver() float64
}


// without pointer receiver
func ( c circle ) area_with_no_pointer_receiver() float64 {
    return math.Pi * c.radius * c.radius
}

// with pointer receiver
func ( c *circle ) area() float64 {
    return math.Pi * c.radius * c.radius
}

// getting the info
func info ( s shape ) {
    fmt.Println("area info", s.area());
}

// create the info_ function to implement a function without receiving the parameter  
func info_ ( s shape ) {
    fmt.Println("area info_", s.area_with_no_pointer_receiver());
}

func main() { 
    c := circle { 
        radius: 5,
    }

    // a pointer receiver will run the function, 
    // pointer to from the circle
    info(&c);

    // a function without the pointer receiver  will run the address 
    info_(&c);

    // this will not run 
    // info(c);
}

