package main

import (
    "fmt"
    "math"
    "sort"
)

var val_list = make(map[int][][]int)
var plist = build_prime_list(100)

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

func build_prime_list(limit int) []int {
    plist := make([]int, 0)
    plist = append(plist, 0) //
    
    for i := 2; i < limit; i++ {
        if is_prime(i) {
            plist = append(plist, i)
        }
    }
    return plist
}

func sum_equals(arr []int, val int) bool {
    //can be sped up
    v := 0
    for _, x := range arr {
        v += x
    }
    return v == val
}

func get_combies(n int) [][]int {
    // return all combinations for one number e.g. 5 = ((2,3))
    //start with worst case, not using any previous values
    
    combination_length := n/2
    index := sort.Search(len(plist), func(i int) bool { return plist[i] >= n })
    all_combos := combrep(combination_length, plist[:index - 1]) //TODO start at n
    found := make([][]int, 0)
    for _, v := range all_combos {
        if sum_equals(v, n) {
            fmt.Println("found", v)
            found = append(found, v)    
        }
    }
    return found
    // fmt.Printf("n(%v) = %v\n", n, prime_sum_count)
}

func main() {
    val_list[1] = nil
    val_list[2] = nil
    val_list[3] = nil
    val_list[4] = append(val_list[4], []int{2,2})
    num_limit := 18
    for i := 5; i <= num_limit; i++ {
        val_list[i] = get_combies(i)
    }

    fmt.Println(len(val_list[18]))
}