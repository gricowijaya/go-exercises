package main

import("fmt")

func main() { 
  m := map[string]int{
    "James":32,
    "Miss Moneypenny":27,
  }
  fmt.Println(m)
  fmt.Println(m["James"])
  fmt.Println(m["Barnabas"])

  v, ok := m["Barnabas"]
  fmt.Println(v)
  fmt.Println(ok)

  m["Todd"] = 44
  if v, ok := m["Barnabas"]; ok { fmt.Println(v) }

  for k, v := range m {
    fmt.Println("Print the ", k, v);
  }

  xi := []int{4,2,3,4,1,54}

  for i, v := range xi {
    fmt.Println(i, v);
  }
}
