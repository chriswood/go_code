package main

import (
    "fmt"
    "math"
    "sort"
    "time"
)

var val_list = make(map[int][][]int)
var plist = build_prime_list(100)

func display() {
    for i := 1; i < len(val_list); i++ {
        fmt.Printf("%v - %v\n", i, len(val_list[i]))
    }
}

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
    //SKIPPING 2, 3, 5
    for i := 7; i < limit; i++ {
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

func comp_array(a []int, b []int) bool {
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func deep_copy(a []int) []int {
    result := make([]int, 0)
    for _, el := range a {
        result = append(result, el)
    }
    return result
}

func exists(a [][]int, b []int) bool {
    temp_a, temp_b := []int{}, []int{}
    for _, element := range a {
        if len(element) == len(b) {
            temp_a = deep_copy(element)
            temp_b = deep_copy(b)
            sort.Ints(temp_a)
            sort.Ints(temp_b)
            if comp_array(temp_a, temp_b) {
                return true
            }
        }
    }
    return false
}

func get_combies(n int) {
    // save all combinations for one number e.g. 5 = ((2,3))
    combination_length := n/7
    index := sort.Search(len(plist), func(i int) bool {return plist[i] >= n - 1})
    all_combos := combrep(combination_length, plist[:index])
    //fmt.Println("all_combos for", n,plist[:index - 1])
    for _, v := range all_combos {
        if sum_equals(v, n) {
            val_list[n] = append(val_list[n], v)
        }
    }
    // fmt.Printf("n(%v) = %v\n", n, prime_sum_count)
}

func main() {
    defer display()
    start := time.Now()
    val_list[1] = nil
    val_list[2] = nil
    val_list[3] = nil
    val_list[4] = append(val_list[4], []int{2,2})
    val_list[5] = append(val_list[5], []int{2,3})
    val_list[6] = append(val_list[6], []int{3,3}, []int{2,2,2})
    val_list[7] = append(val_list[7], []int{2,2,3}, []int{5,2})
    val_list[8] = append(val_list[8], []int{2,3,3}, []int{2,2,2,2}, []int{3,5})
    num_limit := 72

    //first run all larger combinations, no dupes possible

    for i := 9; i <= num_limit; i++ {
        fmt.Println("starting combinations for", i)
        get_combies(i)
    }
    fmt.Printf("combies took in %s\n", time.Now().Sub(start))
    start = time.Now()
    //handle 2,3
    for i := 9; i <= num_limit; i++ {

        //**********2*************
        for _, v := range val_list[i - 2] {
            temp := append(v, 2)
            //fmt.Println("i,temp",i,temp)
            val_list[i] = append(val_list[i], temp)
        }
        //then add new one if needed
        if i % 2 == 1 && is_prime(i - 2) {
            val_list[i] = append(val_list[i], []int{2, i - 2})
        }
        //**********3*************
        for _, v := range val_list[i - 3] {
            temp := append(deep_copy(v), 3)
            if !exists(val_list[i], temp) {
                val_list[i] = append(val_list[i], temp)
            }
        }
        if is_prime(i - 3) {
            val_list[i] = append(val_list[i], []int{3, i - 3})
        }
        //**********5*************
        for _, v := range val_list[i - 5] {
            temp := append(deep_copy(v), 5)
            if !exists(val_list[i], temp) {
                val_list[i] = append(val_list[i], temp)
            }
        }
        if is_prime(i - 5) {
            val_list[i] = append(val_list[i], []int{5, i - 5})
        }
    }

    fmt.Printf("2 handling took %s\n", time.Now().Sub(start))
}