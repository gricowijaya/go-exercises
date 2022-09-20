package main

import("fmt")

func main() {
  slices := []int{1,3,2,4,4,4,5,6,7,1}

  for _, values := range slices {
    fmt.Println(values);
  }

  fmt.Println(`before slices`, slices);

  slices = append(slices[:0], slices[3:]...) // delete from index 0 until 3

  newSlices := []int{8,8,8,876,6}

  slices = append(newSlices);

  fmt.Println(`after slice`, slices);
}


