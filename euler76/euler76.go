package main

import (
    "fmt"
    "math"
)

//p(k) = p(k − 1) + p(k − 2) − p(k − 5) − p(k − 7) + p(k − 12) + p(k − 15) − p(k − 22) − ...
//where p(0) is taken to equal 1, and p(k) is taken to be zero for negative k.

func sign(k int) int {
    exp := math.Floor(float64((k + 3)/2))
    return int(math.Pow((-1), exp))
}

func main(){
    for i := 1; i < 10; i++ {
        fmt.Println(i, sign(i))
    }
}