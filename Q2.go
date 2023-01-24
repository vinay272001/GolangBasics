// Initialize the constants using iota and you should find the correct formula for iota
//             for printing months 11,10,9 
//               const(
//                   Nov = 11
//                   Oct = 10
//                   Sep = 9
//                 )
//                print(Nov,Oct,Sep)  
package main

import "fmt"

const (
    Nov = 11 - iota
    Oct
    Sep
)

func main() {
    fmt.Println(Nov, Oct, Sep)
}
