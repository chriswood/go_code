package main

import (
    "fmt"
    "math"
)

func is_prime(n int) bool {
    // check quick values
    if n == 4 {
        return False
    }
    if n < 4 {
        return True
    }

    //get last value to test for even division into N
    last_value := 4//math.Ceil(math.Sqrt(n))

    for i := 2; i <= last_value; i++ {
        if n % i == 0 {
            return False
        }
    }
    return True
}

func main() {
    p := is_prime(8)
    fmt.println(p)
}

