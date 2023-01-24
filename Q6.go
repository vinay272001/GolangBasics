// Print statements
// Print your own names (eg : “abc”) using different print statements i.e Print(), Println(), Printf() 
// Output should be like this
// I am abc
// I am abc
// I am abc
package main

import "fmt"

func main() {
    name := "abc"
    fmt.Print("I am ")
    fmt.Print(name)
    fmt.Println() // new line

    fmt.Println("I am", name)

    fmt.Printf("I am %s\n", name)
}
