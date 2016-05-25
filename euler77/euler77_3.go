package main

import (
    "fmt"
    "math"
)

func upper_middle(x int) int {
    a := float64(x)
    high_float := math.Ceil(math.Sqrt(a))
    return int(high_float)
}

func is_prime(n int) bool {
    // check quick values
    if n == 4 {
        return false
    } else if n < 4 {
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

func build_prime_list(stop int) []int {
    plist := make([]int, 0)
    //plist = append(plist, 0)
    for i := 2; i < (stop - 1); i++ {
        if is_prime(i) {
            plist = append(plist, i)
        }
    }
    return plist
}

func main() {
    //array[sum total] = ways to make it
    //for each prime p:
    //    for each prime t > p:
    limit := 13 //maximum subtotal we compute
    plist := build_prime_list(100)
    partitions := make(map[int]int)
    partitions[0] = 0
    partitions[1] = 0
    p1,p2 := 0,0
    for i := 0; p1 < limit; i++ {
        p1 = plist[i] //2
        for j := i; p2 < limit; j++ {
            p2 = plist[j]
            fmt.Println("p1,p2",p1,p2)
            partitions[p1 + p2]++
        }
        p2 = 0
        fmt.Println(p1)
    }
    fmt.Println(partitions)
}
// p = 2,3
// partitions[2] = 1 = how many ways to suym to 2
// partitions[3] = 0
// partitions[4] = 1
// partitions[5] = 0
// partitions[6] = 1
// partitions[7] = 0
// 9 will have p(7) + 2
// and p(3) + 3
