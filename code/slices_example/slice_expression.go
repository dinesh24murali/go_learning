package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5, 6}
    subSlice := slice[1:4]
    fmt.Println(subSlice)                     // => [2 3 4]
    fmt.Println(len(subSlice), cap(subSlice)) // => 3 3

    subSliceWithCap := slice[1:4:5]
    fmt.Println(subSliceWithCap)                            // => [2 3 4]
    fmt.Println(len(subSliceWithCap), cap(subSliceWithCap)) // => 3 4
}