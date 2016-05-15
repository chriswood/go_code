package main

import (
    "fmt"
    "math"
    "time"
)

func gcd(a int, b int) int {
        //return a if b == 0 else gcd(b, a % b)
    if b == 0 {
        return a
    } else {
        return gcd(b, a % b)
    }
}

func single_triplet(s int) bool {
    k, subtotal := 0, 0
    s2 := s/2
    mlimit := int(math.Ceil(math.Sqrt(float64(s2)))) - 1
    for m := 2; m <= mlimit; m++ {
        if s2 % m == 0 {
            sm := s2/m
            for sm % 2 == 0 {
                sm = sm/2
            }

            if m % 2 == 1 {
                k = m + 2
            } else {
                k = m + 1
            }
            for k < 2*m && k <= sm {
                if sm % k == 0 && gcd(k,m) == 1 {
                    subtotal++
                }
                if subtotal > 1 { //save a few loops
                    return false
                }
                k += 2
            }
        }
    }
    return subtotal == 1
}

func check(s int) int {
    //a2+b2=c2
    tally := 0
    for a := 3; a <= (s-3)/3; a++ {
        for b := (a+1); b <= (s - 1 - a)/2; b++ {
            c := s-a-b
            if c*c == a*a + b*b{
                tally++
            }
        }
    }
    return tally
}

func main() {
    start := time.Now()
    total := 0
    limit := 1500000
    //problem with 108
    for i := 12; i <= limit; i += 2 {
        if i % 10000 == 0 {
            fmt.Println(i)
        }
        if single_triplet(i) == true {
            total++
        }
    }
    fmt.Println("There were ", total, "triangles")
    end := time.Now()
    fmt.Printf("Calculation finished in %s\n", end.Sub(start))
}