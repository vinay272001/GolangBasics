// If, if-else ,if-else-if statement
//             Print expected output based on age variable
//             Expected output : 
//                 If age is greater than 60 print : Getting older
//                 If age is greater than 30 print: Getting wiser
//                 If age is greater than 20 print: Adulthood
//                 If age is greater than 10 print: Younger
package main

import "fmt"

func main() {
    age := 25

    if age > 60 {
        fmt.Println("Getting older")
    } else if age > 30 {
        fmt.Println("Getting wiser")
    } else if age > 20 {
        fmt.Println("Adulthood")
    } else if age > 10 {
        fmt.Println("Younger")
    }
}
