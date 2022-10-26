package benchmarking

import (
    "testing"
    "fmt"
)

func TestGreet(t *testing.T) { 
    s := Greet("James")
       if s != "Welome James" { 
            t.Error("got", s, "expected", "Welcome James")
       }
}

func ExampleGreet() { 
    fmt.Println(Greet("James"))
    // Output:
    // Welcome James
}

func BenchmarkGreet(b *testing.B) { 
    // benchmarking will run your code a lot of times that's why you should use the loop, so it'll determine how many time it runs.
    for i:= 0; i< b.N; i++ { 
        Greet("James")
    }
}
