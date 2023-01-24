// Functions
// Create a function for multiplying 3 numbers and call that function in the the main function and print the result
package main

import "fmt"

func multiply(a, b, c int) int {
    return a * b * c
}

func main() {
    result := multiply(2, 3, 4)
    fmt.Println("Result:", result)
}
