package main

import (
    "fmt"
    "time"
)

var limit int = 1000000

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

func calc_phi(n int) int {
    sub_total := 1
    for i := 2; i < n; i++ {
        if is_coprime(i, n) {
            sub_total++
        }
    }
    return sub_total
}

func main() {
    start := time.Now()
    total := 0
    for i := 1; i <= limit; i++ {
        if i % 1000 == 0 {
            fmt.Println(i)
        }
        total += calc_phi(i)
    }
    fmt.Println("fractions", total - 1)
    elapsed := time.Since(start)
    fmt.Printf("This took %s", elapsed)
    fmt.Println(" ")
}
