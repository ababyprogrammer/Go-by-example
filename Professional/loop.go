package main

import "fmt"

func main() {
    squares := make([]int, 5) // creates an array with 5 elements, all initially zero

    for i := 0; i < 5; i++ {
        squares[i] = i * i // assigns the square of i to the i-th element of the array
    }

    fmt.Println(squares) // prints the array of square numbers
}