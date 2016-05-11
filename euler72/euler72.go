package main

import (
    "fmt"
    "math"
    "time"
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
            }
            //first two divisors are prime, we're done with this iteration
            return
        }
    }
}

func phi(phi float64) float64 {
    encountered := map[int]bool{}
    for _, factor := range plist {
        if encountered[factor] == false {
            phi *= (1 - 1.0/float64(factor))
            encountered[factor] = true
        }
    }
    return phi
}

func main() {
    start := time.Now()
    total := 0.0
    limit := 1000000
    for i := 2; i <= limit; i++ {
        if i % 10000 == 0 {
            fmt.Println(i)
        }
        plist = nil
        prime_factors(i)
        total += phi(float64(i))
    }
    fmt.Println("total was", total)
    elapsed := time.Since(start)
    fmt.Printf("This took %s", elapsed)
    fmt.Println(" ")
}
