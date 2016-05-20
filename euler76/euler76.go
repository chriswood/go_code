package main

import (
    "fmt"
    "math"
    "time"
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
var pent_map = make(map[int]int)
var partition_terms = make(map[int]int)

func get_sign(k int) int {
    exp := math.Floor(float64((k)/2))
    return int(math.Pow((-1), exp))
}

func pentagonal_gen(c chan int, n int) {
    //Send both positive and negative "general" values to channel
    //n(3n-1)/2
    p1, p2 := 0, 0
    for i := 1; i <= n; i++ {
        anti_i := (-1) * i
        p1 = (i*(3*i - 1))/2
        p2 = (anti_i*(3*anti_i - 1))/2
        c <- p1
        c <- p2
    }
    close(c)
}


func get_partition_term(n int) int {
    if n < 0 {
        return 0
    } else if n == 0 {
        return 1
    } else {
        return partition_terms[n]
    }
}

func build_term_controller(i int) int {
    //i is the term we want the partition sequence for, p(i)
    //build the proper sequence for just that term, and return the value
    pval := 0
    for j := 0; j <= i; j++ {
        pval += get_sign(j) * get_partition_term(i - pent_map[j])
    }
    return pval
}

func build_pentagonal_list(limit int) {
    //no need to do this more than once
    c := make(chan int)
    go pentagonal_gen(c, limit)
    for i := 0; i <= limit; i++ {
        pent_map[i] = <-c
    }
}

func main() {
    start := time.Now()
    limit := 11
    build_pentagonal_list(limit)
    partition_terms[0] = 1
    for i := 1; i <= limit; i++ {
        partition_terms[i] = build_term_controller(i)
    }
    //Note the answer for the problem is one less than the partition function value
    fmt.Println("p(",limit,") = ",partition_terms[limit] - 1)
    end := time.Now()
    fmt.Printf("Calculation finished in %s\n", end.Sub(start)) 
}