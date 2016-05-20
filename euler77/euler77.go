package main

//What is the first value which can be written as the sum of primes in over five thousand different ways?
import (
    "fmt"
    "math"
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

func build_prime_list(limit int) []int {
    plist := make([]int, 0)
    count := 0
    for i := 2; count < limit; i++ {
        if is_prime(i) {
            plist = append(plist, i)
            count++
        }
    }
    return plist
}

func combrep(n int, clist []int) [][]int {
    if n == 0 {
        return [][]int{nil}
    }
    if len(clist) == 0 {
        return nil
    }
    r := combrep(n, clist[1:])
    for _, x := range combrep(n-1, clist) {
        r = append(r, append(x, clist[0]))
    }
    return r
}

func main() {
    test_val := 3
    plist := build_prime_list(test_val)
    //test_arr := []int{1, 2, 3}
    fmt.Println(combrep(2, plist))
}



