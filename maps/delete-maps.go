package main

import("fmt")

func main() { 
  m := map[string]int{
    "James":32,
    "Miss Moneypenny":27,
  }
  fmt.Println(m)
   delete(m, "James")

   delete(m, "Ian Fleming")

  fmt.Println(m["Miss Moneypenny"])
  fmt.Println(m["Ian Fleming"])

  if v, ok := m["Miss Moneypenny"]; ok {
    fmt.Println("value", v);
    delete(m, "Miss moneypennny")
  }
}
