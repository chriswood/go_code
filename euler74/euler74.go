
package main

import (
    "fmt"
    "time"
    "strconv"
)
const LIMIT int = 15
var facts = make(map[int]int)

func build_map() {
    facts[0] = 1
    facts[1] = 1
    n := 1
    for i := 2; i < 10; i++ {
        n *= i
        facts[i] = n
    }
}

func factorial_sum(n int) int {
    s := strconv.Itoa(n)
    total := 0
    digit := 0
    for _, d := range s {
        digit = int(d - '0')
        total += facts[digit]
    }
    return total
}

func chain_length(n int) int {
    count := 0
    m := map[int]bool {n: true}
    for {
        sum := factorial_sum(n)
        count++
        _, exists := m[sum]
        if exists {
            return count
        }
        m[sum] = true
        n = sum
    }
}

func main() {
    start := time.Now()
    build_map()
    total := 0

    for i := 1; i < 1000000; i++ {
        if i % 5000 == 0 {
            fmt.Println(i)
        }
        result := chain_length(i)
        if result == 60 {
            total++
        }
    }
    fmt.Println("There were ", total, "terms")
    end := time.Now()
    fmt.Printf("Calculation finished in %s\n", end.Sub(start)) // Calculation finished in 0ms
}

