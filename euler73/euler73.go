package main

import (
    "fmt"
)

func gcd(a int, b int) int {
        //return a if b == 0 else gcd(b, a % b)
    if b == 0 {
        return a
    } else {
        return gcd(b, a % b)
    }
}

func is_coprime(i int, j int) bool {
    //return False if (b % 2 == 0) and (a % 2 == 0) else gcd(a, b) == 1
    if i % 2 == 0 && j % 2 == 0 {
        return false
    }
    return gcd(i, j) == 1
}

func in_range(a int, b int) bool {
    af, bf := float64(a), float64(b)
    if 1.0/3.0 < af/bf && af/bf < 1.0/2.0 {
        return true
    }
    return false
}

func get_fractions_count(n int) int {
    count := 0 //starting at 2
    for i := 2; i < n; i++ {
        if is_coprime(i, n) {
            if in_range(i, n) {
                //fmt.Println("i,n", i, n)
                count++
            }
        }
    }
    return count
}

func main() {
    const limit int = 12000
    fcount := 0
    for i := 2; i <= limit; i++ {
        if i % 1000 == 0 {
            fmt.Println(i)
        }
        fcount += get_fractions_count(i)
    }
    fmt.Println("count is", fcount)
}