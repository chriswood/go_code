package main

//What is the first value which can be written as the sum of primes in over five thousand different ways?
import (
    "fmt"
    "math"
    "time"
)

var val_list = make(map[int]int)

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
    plist = append(plist, 0) //
    
    for i := 3; i < limit; i++ {
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
            //fmt.Println("found", v)
            total++
        }
    }
    return total
}

func main() {//2,2,2,2,3 - 2,2,2,5 - 2,2,7 - 3,3,5 - 3,3,3,2
    ////9 == 4,10 = 5, 11=5,12 = 7, 13=8, 14 = 10, 15=12, 16 = 14, 17=16, 18 = 19, 19=22,20 = 26
    val_list[8] = 3
    val_list[9] = 4
    val_list[10] = 5
    val_list[11] = 5
    n := 12
    prime_sum_count := 0
    limit := 30

    for prime_sum_count < limit {
        start := time.Now()
        combination_length := n/2
        plist := build_prime_list(n - 1)
        all_combos := combrep(combination_length, plist)
        prime_sum_count = val_list[n - 2] + sum_count(all_combos, n)
        if n % 2 == 1 && is_prime(n - 2) { //handle 2 plus n - 2
            prime_sum_count++
        }
        val_list[n] = prime_sum_count
        fmt.Printf("finished in %s\n", time.Now().Sub(start))
        fmt.Printf("n(%v) = %v\n", n, prime_sum_count)
        n++
    }
    //fmt.Printf("%v\n", val_list)
}

//12 all the ways 9 can be summed by 3 plus 1



