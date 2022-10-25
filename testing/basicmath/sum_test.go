package basicmath

import "testing"

// First Test
// func TestSum(t *testing.T) { 
//     // here we create the result variable
//     var result int 
//     // let's put the return value of sum(a, b) into result
//     result = Sum(1, 3)
//     // if 1 + 3 is not equal to 4 then print error
//     if result != 4 { 
//         t.Error("Expected 4, got ", result)
//     }
// }

// Second Test

type testpair struct { 
    // need to be a slice of int because we will assign a value for parameter a and b
    values []int
    // the returned result of Sum()
    results int
}

// Let's create an array of testpair.
var tests = []testpair { 
    // values , results
    {[]int{1,3}, 4},
    {[]int{2, 4}, 6},
    {[]int{10, 4}, 14},
    {[]int{10, 134}, 144},
}

// refactor
func TestSum(t *testing.T) { 
    // loop through the tests value
    for _, pair := range tests { 
        // in the function of Sum we pass two parameters then we will get the value of index [0] and index[1] of the slice of values
        resultValue := Sum(pair.values[0], pair.values[1])
         // check the result
         if resultValue != pair.results { 
            // If there's an error give the error 
            t.Error(
             "For ", pair.values, 
             "expected", pair.results, 
             "got", resultValue, 
            )
         } 
    }
}

