package main

//What is the first value which can be written as the sum of primes in over five thousand different ways?
import (
    "fmt"
    "math"
    "time"
)


func upper_middle(x int) int {
    // returns the integer ceil of x/2
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

func build_prime_list(limit int, start int) []int {
    plist := make([]int, 0)
    plist = append(plist, 0) //
    for i := start; i < limit; i++ {
        if is_prime(i) {
            plist = append(plist, i)
        }
    }
    return plist
}

func combrep(n int, clist []int) [][]int {
    //fmt.Println(n,"---", clist)
    if n == 0 {
        return [][]int{nil}
    }
    if len(clist) == 0 {
        return nil
    }
    r := combrep(n, clist[1:])
    a := make([]int,0)
    for _, x := range combrep(n - 1, clist) {
        a = append(x, clist[0])
        r = append(r, a)
    }
    return r
}

func sum_equals(arr []int, val int) bool {
    //can be sped up
    v := 0
    for _, x := range arr {
        v += x
    }
    return v == val
}

func sum_count(whole_arr [][]int, val int) int {
    total := 0
    for _, v := range whole_arr {
        if sum_equals(v, val) {
            fmt.Println("found", v)
            total++
        }
    }
    return total
}

func main() {
    //9 == 4,10 = 5, 11=5,12 = 7, 13=8, 14 = 10, 15=12, 16 = 14, 17=16, 18 = 19, 19=22,20 = 26
    start := time.Now()
    n := 17
    plist := build_prime_list(n - 1, 2)
    combination_length := n/2
    fmt.Println(plist)
    all_combos := combrep(combination_length, plist)
    //fmt.Println(all_combos)
    prime_sum_count := sum_count(all_combos, n)
    fmt.Printf("finished in %s\n", time.Now().Sub(start))
    fmt.Printf("n(%v) = %v\n", n, prime_sum_count)
}