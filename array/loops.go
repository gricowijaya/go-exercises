package main

import("fmt")

func main() {
  var number_of_array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  for i := 0; i <=9 ; i++ {
    fmt.Println(`using the for loop`, number_of_array[i]);
  }

  for _, value := range number_of_array {       // use a blank identifiers 
    fmt.Println(`using the for in range loop `, value);
  }
}
