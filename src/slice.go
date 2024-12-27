package main

import "fmt"

func main() {
    // Create a slice
    slice := []int{1, 2, 3, 4, 5}

    // Access and modify elements
    fmt.Println("Original slice:", slice)
    slice[2] = 10
    fmt.Println("Modified slice:", slice)

    // Append elements
    slice = append(slice, 6, 7)
    fmt.Println("After appending:", slice)

    // Length and capacity
    fmt.Println("Length:", len(slice))
    fmt.Println("Capacity:", cap(slice))

    // Slice a slice
    subSlice := slice[1:4]
    fmt.Println("Sub-slice:", subSlice)

    // Copy a slice
    copied := make([]int, len(slice))
    copy(copied, slice)
    fmt.Println("Copied slice:", copied)
}
