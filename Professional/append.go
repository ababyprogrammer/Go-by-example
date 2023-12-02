package main

import "fmt"

func main() {
    squares := make([]int, 0, 5) // creates an empty slice with a capacity of 5

    for i := 0; i < 5; i++ {
        squares = append(squares, i*i) // appends the square of i to the slice
    }

    fmt.Println(squares) // prints the slice of square numbers
}