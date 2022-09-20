/* * create a type square */
/* * create a type circle */
/* * attach a method to each that calculates area and returns it */
/* * create a type shape which defines an interface as anything which */ 
/* has the area method */
/* * create a func info which takes type shape and then prints the area. */
/* * create a value of type square */
/* * create a value of type circlek */
/* * use func info to print the area of square */
/* * use func info to print the area of circle */


// This is the first attempt
package main

// instead we can use the "math" package
import (
  "fmt"
  "math"
)

// instead of doing this use the package math
const PI float64 = 3.14

// interface for counting area so it'll do polymorphism
type object interface {
  area() float64
}

type square struct {
  side int
}

type circle struct {
  radius float64
}

func(s square) area() float64 {
  var result int
  result = s.side * s.side
  return float64(result)
}

func(c circle) area() float64 {
  // the results should be float64
  var result float64 
  // result = PI * float64(c.radius) * float64(c.radius)
  result = math.Pi * float64(c.radius) * float64(c.radius) // v2
  return result
}

// we also can use the shape function just to print the area such as below
// we have create the object interface and now we can do the print area 
func info(o object) {
  fmt.Println(o.area());
}


func main() {
  square := square {33,}
  circle := circle {4,}

  fmt.Printf("%.2f is the area of circle from radius %d\n", circle.area(), int(circle.radius));
  fmt.Printf("%d is the area of square from side %d\n", int(square.area()), square.side);

  // output with the method of info
  fmt.Println("Printed with info method")
  info(circle)
  info(square)

}

