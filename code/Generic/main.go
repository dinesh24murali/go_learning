package main

import "fmt"
import "golang.org/x/exp/constraints"


// normal function without Generic
func Min(x, y float64) float64 {
    if x < y {
        return x
    }
    return y
}

func main() {
	var i float64 = Min(12.4,13.4)
	fmt.Println(i)
}



// Function with generic

func GMin[T constraints.Ordered](x, y T) T {
    if x < y {
        return x
    }
    return y
}
