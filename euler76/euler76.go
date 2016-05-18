package main

import (
    "fmt"
    "math"
)

//p(k) = p(k − 1) + p(k − 2) − p(k − 5) − p(k − 7) + p(k − 12) + p(k − 15) − p(k − 22) − ...
//where p(0) is taken to equal 1, and p(k) is taken to be zero for negative k.
// p0 = 1
// p1 = 1
// p2 = p1 + p0 = 2
// p3 = p2 + p1 = 3
// p4 = p3 + p2 = 5
// p5 = p4 + p3 - p0 = 7
// p6 = p5 + p4 - p1 = 11
// p7 = p6 + p5 - p2 - p0 = 15

func sign(k int) int {
    exp := math.Floor(float64((k + 3)/2))
    return int(math.Pow((-1), exp))
}

func pentagonal_gen(c chan int, n int) {
    //n(3n-1)/2
    p1,p2 := 0,0
    for i := 1; i <= n; i++ {
        anti_i := (-1) * i
        p1 = (i*(3*i - 1))/2
        p2 = (anti_i*(3*anti_i - 1))/2
        c <- p1
        c <- p2
    }
    close(c)
}


func partition(n int) int {
    if n < 0 {
        return 0
    } else if n == 0 {
        return 1
    } else {
        return 8//FIX
    }
}

func build_term(k int, x int) int {
    fmt.Println("x",x)
    return x
}

func main() {
    limit := 10
    terms := make(map[int]int)
    c := make(chan int)
    terms[0] = 1
    go pentagonal_gen(c, limit)
    for i := 1; i <= limit; i++ {
        terms[i] = build_term(i, <-c)
    }
    // for pval := range c {
    //     terms[]
    // }
}