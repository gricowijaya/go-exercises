Starting with this code, marshal the []user to JSON.

```go
package main

import "fmt"

type user struct {
    first string
    age int
} 

func main() {
    u1 := user {
        first: "Dani",
        age: 18,
    }

    u2 := user {
        first: "Rama",
        age: 38,
    }

    u3 := user {
        first: "Ajeng",
        age: 20,
    }

    users := []user{u1, u2, u3} 
    
    // here
}
```

