package main

import (
    "fmt"
    "math"
    "os"
)

var plist = make([]int, 0)
var count int = 0

func upper_middle(x int) int {
    // returns the ceil of x/2
    a := float64(x)
    high_float := math.Ceil(math.Sqrt(a))
    return int(high_float)
}

func is_prime(n int) bool {
    // check quick values
    if n == 4 {
        return false
    }
    if n < 4 {
        return true
    }

    last_value := upper_middle(n)

    for i := 2; i <= last_value; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func prime_factors(n int) {
    if count > 3 {
        os.Exit(3)
    } else {
        count++
    }
    fmt.Println("being ran with ", n)
    fmt.Println("plist = ", plist)
    //First, test for primeness
    if is_prime(n) {
        plist = append(plist, n)
        return
    }
    end := int(math.Ceil(float64(n)/2.0))
    for i := 2; i <= end; i++ {
        if n % i == 0 {
            if is_prime(i) {
                plist = append(plist, i)
            }
            factor2 := n/i
            if is_prime(factor2) {
                plist = append(plist, int(factor2))
            } else {
               prime_factors(factor2)
               break
            }
        }
    }
    return
}

func main() {
    //p := prime_factors(17)

    prime_factors(72)
    fmt.Println(plist)
}
